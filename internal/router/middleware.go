package router

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ad9311/renio-go/internal/action"
	"github.com/ad9311/renio-go/internal/envs"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/vars"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
)

type jwtClaims struct {
	JTI string
	SUB int
}

var freeRoutes = []string{
	"/info",
	"/auth/sign-in",
	"/auth/sign-up",
}

func headerRouter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")
		next.ServeHTTP(w, r)
	})
}

func routesProtector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if findInFreeRoutes(path) {
			next.ServeHTTP(w, r)
			return
		}

		errResponse := action.ErrorResponse{}
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			errResponse.Append(fmt.Errorf("invalid request authorization"))
			action.WriteError(w, errResponse)
			return
		}

		splitValue := strings.Split(authHeader, " ")
		if len(splitValue) != 2 {
			errResponse.Append(fmt.Errorf("invalid request authorization"))
			action.WriteError(w, errResponse)
			return
		}

		jwtToken := splitValue[1]
		claims, err := decodeJWT(jwtToken)
		if err != nil {
			errResponse.Append(fmt.Errorf("invalid jwt token"))
			action.WriteError(w, errResponse)
			return
		}

		var allowedJWT model.AllowedJWT
		err = allowedJWT.SelectByJTI(claims.JTI)
		if err == pgx.ErrNoRows {
			errResponse.Status = http.StatusUnauthorized
			errResponse.Append(fmt.Errorf("not authorized"))
			action.WriteError(w, errResponse)
			return
		}

		if allowedJWT.UserID != claims.SUB {
			errResponse.Append(fmt.Errorf("%d - %d", allowedJWT.UserID, claims.SUB))
			action.WriteError(w, errResponse)
			return
		}

		ctx := context.WithValue(r.Context(), vars.AllowedJWTKey, allowedJWT)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func findInFreeRoutes(path string) bool {
	for _, str := range freeRoutes {
		if str == path {
			return true
		}
	}

	return false
}

func decodeJWT(tokenStr string) (jwtClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(envs.GetEnvs().JWTToken), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		jit := fmt.Sprintf("%s", claims["jti"])
		subStr, _ := claims.GetSubject()
		sub, _ := strconv.Atoi(subStr)

		return jwtClaims{JTI: jit, SUB: sub}, nil
	}

	return jwtClaims{}, err
}
