package router

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/action"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RoutesHandler() http.Handler {
	r := chi.NewRouter()

	// --- Middleware --- //
	r.Use(middleware.Logger)
	r.Use(routesProtector)

	r.Route("/", func(r chi.Router) {
		// --- Auth --- //
		r.Route("/auth", func(r chi.Router) {
		})

		// --- Budget --- //
		r.Route("/budgets", func(r chi.Router) {
			r.Use(action.BudgetAccountCTX)
			r.Route("/", func(r chi.Router) {
				r.Route("/{budgetUID}", func(r chi.Router) {
					r.Use(action.BudgetCTX)

					// --- Income --- //
					r.Route("/income-list", func(r chi.Router) {
						r.Route("/{incomeID}", func(r chi.Router) {
							r.Use(action.IncomeCTX)
						})
					})

					// --- Expense --- //
					r.Route("/expenses", func(r chi.Router) {
						r.Route("/{expenseID}", func(r chi.Router) {
							r.Use(action.ExpenseCTX)
						})
					})
				})
			})
		})
	})

	return r
}
