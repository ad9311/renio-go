package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ad9311/renio-go/internal/action"
	"github.com/ad9311/renio-go/internal/conf"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/golang-jwt/jwt/v5"
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

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			action.WriteError(w, []string{"invalid request authorization"}, http.StatusUnauthorized)
			return
		}

		splitValue := strings.Split(authHeader, " ")
		if len(splitValue) != 2 {
			action.WriteError(w, []string{"invalid request authorization"}, http.StatusUnauthorized)
			return
		}

		jwtToken := splitValue[1]
		claims, err := decodeJWT(jwtToken)
		if err != nil {
			action.WriteError(w, []string{"invalid jwt token"}, http.StatusUnauthorized)
			return
		}

		var allowedJWT model.AllowedJWT
		err = allowedJWT.FindByJTI(claims.JTI)
		if err != nil {
			action.WriteError(w, []string{err.Error()}, http.StatusUnauthorized)
			return
		}
		if allowedJWT.UserID != claims.SUB {
			str := fmt.Sprintf("%d - %d", allowedJWT.UserID, claims.SUB)
			action.WriteError(w, []string{str}, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), conf.UserIDContext, allowedJWT.UserID)
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

		return []byte(os.Getenv("JWT_KEY")), nil
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
