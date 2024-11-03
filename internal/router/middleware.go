package router

import (
	"net/http"
	"strings"

	"github.com/ad9311/renio-go/internal/conf"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/vars"
)

var freeRoutes = []string{
	"/auth/sign-in",
	"/auth/sign-up",
}

func routesProtector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if findInPublicRoutes(path) {
			next.ServeHTTP(w, r)
			return
		}

		_, ok := conf.GetSession().Get(r.Context(), string(vars.CurrentUserKey)).(model.User)
		if !ok {
			http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func findInPublicRoutes(path string) bool {
	for _, str := range freeRoutes {
		if str == path {
			return true
		}
	}

	return strings.HasPrefix(path, "/static/")
}
