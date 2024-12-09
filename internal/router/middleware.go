package router

import (
	"context"
	"net/http"
	"strings"

	"github.com/ad9311/renio-go/internal/app"
	"github.com/ad9311/renio-go/internal/handler"
	"github.com/ad9311/renio-go/internal/vars"
	"github.com/justinas/nosurf"
)

func session(next http.Handler) http.Handler {
	return app.GetSession().LoadAndSave(next)
}

func csrf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.GetEnv().AppEnv == app.Production,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if isPublicResource(path) {
			next.ServeHTTP(w, r)
			return
		}

		ctx := r.Context()

		isUserSignedIn := handler.IsUserSignedInCTX(ctx)
		isSignInPath := strings.HasPrefix(path, "/auth/sign-in")
		isSignUpPath := strings.HasPrefix(path, "/auth/sign-up")

		if (isSignInPath || isSignUpPath) && isUserSignedIn {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		if isSignUpPath && !isUserSignedIn {
			next.ServeHTTP(w, r)
			return
		}

		if !strings.HasPrefix(path, "/auth/sign-in") && !isUserSignedIn {
			http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func appData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := handler.GetUserCTX(ctx)
		isUserSignedIn := handler.IsUserSignedInCTX(ctx)

		data := handler.TmplData{
			"errors":         []string{},
			"alert":          "",
			"info":           "",
			"currentUser":    user,
			"isUserSignedIn": isUserSignedIn,
			"csrfToken":      nosurf.Token(r),
			"navLinks":       handler.GetNavLinks(ctx),
			"appEnv":         app.GetEnv().AppEnv,
		}

		ctx = context.WithValue(ctx, vars.AppDataKey, data)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// --- Helpers --- //

func isPublicResource(path string) bool {
	static := strings.HasPrefix(path, "/static/")
	favicon := path == "favicon.ico"

	return static || favicon
}
