package router

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/ctrl"
	"github.com/ad9311/renio-go/internal/midware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RoutesHandler() http.Handler {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(midware.HeaderRouter)
	r.Use(midware.RoutesProtector)

	r.Route("/", func(r chi.Router) {
		// Info
		r.Route("/info", ctrl.InfoRouter(r))

		// Auth
		r.Route("/auth", func(r chi.Router) {
			// Sessions
			r.Route("/sign-in", ctrl.SignInRouter(r))

			// Sign Up
			r.Route("/sign-up", ctrl.SignUpRouter(r))
		})

		// Budgets
		r.Route("/budgets", ctrl.BudgetRouter(r))
	})

	return r
}
