package mw

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ad9311/renio-go/internal/ct"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

type CurrentUserContext string

const UserIDContext = CurrentUserContext("currentUserID")

// Middlewares //

func HeaderRouter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func RoutesProtector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/auth/sign-in" || path == "/info" {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			ct.WriteError(w, []string{"not authorized"}, http.StatusUnauthorized)
			return
		}

		splitValue := strings.Split(authHeader, " ")
		if len(splitValue) != 2 {
			ct.WriteError(w, []string{"not authorized"}, http.StatusUnauthorized)
			return
		}

		jwtToken := splitValue[1]
		jti, err := decodeJWT(jwtToken)
		if err != nil {
			ct.WriteError(w, []string{err.Error()}, http.StatusUnauthorized)
			return
		}

		var allowedJWT model.AllowedJWT
		if err = allowedJWT.FindByJTI(jti); err != nil {
			ct.WriteError(w, []string{err.Error()}, http.StatusUnauthorized)
		}

		ctx := context.WithValue(r.Context(), UserIDContext, allowedJWT.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Helpers //

func decodeJWT(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return fmt.Sprintf("%s", claims["jti"]), err
	}

	return "", err
}
