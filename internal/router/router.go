package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RoutesHandler() http.Handler {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(headerRouter)
	r.Use(routesProtector)

	r.Route("/", func(r chi.Router) {
		// Info
		// r.Route("/info", action.InfoRouter(r))

		// // Auth
		// r.Route("/auth", func(r chi.Router) {
		// 	// Sessions
		// 	r.Route("/sign-in", action.SignInRouter(r))

		// 	// Sign Up
		// 	r.Route("/sign-up", action.SignUpRouter(r))
		// })

		// // Budgets
		// r.Route("/budgets", action.BudgetRouter(r))
	})

	return r
}
