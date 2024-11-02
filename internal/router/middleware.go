package router

import (
	"net/http"
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

func routesProtector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if findInFreeRoutes(path) {
			next.ServeHTTP(w, r)
			return
		}

		next.ServeHTTP(w, r)
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
