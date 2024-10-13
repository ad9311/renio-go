package router

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/ct"
	"github.com/ad9311/renio-go/internal/mw"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RoutesHandler() http.Handler {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(mw.HeaderRouter)
	r.Use(mw.RoutesProtector)

	r.Route("/", func(r chi.Router) {
		// Info
		r.Route("/info", ct.InfoRouter(r))

		// Auth
		r.Route("/auth", func(r chi.Router) {
			// Sessions
			r.Route("/sign-in", ct.SignInRouter(r))

			// Sign Up
			r.Route("/sign-up", ct.SignUpRouter(r))
		})
	})

	return r
}
