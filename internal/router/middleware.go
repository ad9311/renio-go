package router

import (
	"context"
	"net/http"
	"strings"

	"github.com/ad9311/renio-go/internal/conf"
	"github.com/ad9311/renio-go/internal/handler"
	"github.com/ad9311/renio-go/internal/vars"
	"github.com/justinas/nosurf"
)

func session(next http.Handler) http.Handler {
	return conf.GetSession().LoadAndSave(next)
}

func csrf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   conf.GetEnv().AppEnv == conf.Production,
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

		session := conf.GetSession()
		key := string(vars.UserSignedInKey)
		isUserSignedIn := session.GetBool(r.Context(), key)

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
		userKey := string(vars.CurrentUserKey)
		isUserSignedInKey := string(vars.UserSignedInKey)

		user := conf.GetSession().Get(r.Context(), userKey)
		isUserSignedIn := conf.GetSession().GetBool(r.Context(), isUserSignedInKey)

		data := handler.TmplData{
			"errors":         []string{},
			"notice":         "",
			"info":           "",
			"currentUser":    user,
			"isUserSignedIn": isUserSignedIn,
			"csrfToken":      nosurf.Token(r),
			"appEnv":         conf.GetEnv().AppEnv,
		}

		ctx := context.WithValue(r.Context(), vars.AppDataKey, data)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// --- Helpers --- //

func isPublicResource(path string) bool {
	static := strings.HasPrefix(path, "/static/")
	favicon := path == "favicon.ico"

	return static || favicon
}
