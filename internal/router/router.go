package router

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RoutesHandler() http.Handler {
	r := chi.NewRouter()

	// --- Middleware --- //
	r.Use(middleware.Logger)
	r.Use(routesProtector)

	fileServer := http.FileServer(http.Dir("./web/static/"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	r.Route("/", func(r chi.Router) {
		r.Get("/home", handler.Root)
		// --- Auth --- //
		r.Route("/auth", func(r chi.Router) {
		})

		// --- Budget --- //
		r.Route("/budgets", func(r chi.Router) {
			r.Use(handler.BudgetAccountCTX)
			r.Route("/", func(r chi.Router) {
				r.Route("/{budgetUID}", func(r chi.Router) {
					r.Use(handler.BudgetCTX)

					// --- Income --- //
					r.Route("/income-list", func(r chi.Router) {
						r.Route("/{incomeID}", func(r chi.Router) {
							r.Use(handler.IncomeCTX)
						})
					})

					// --- Expense --- //
					r.Route("/expenses", func(r chi.Router) {
						r.Route("/{expenseID}", func(r chi.Router) {
							r.Use(handler.ExpenseCTX)
						})
					})
				})
			})
		})
	})

	return r
}
