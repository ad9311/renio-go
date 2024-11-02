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
	r.Use(headerRouter)
	r.Use(routesProtector)

	r.Route("/", func(r chi.Router) {
		// --- Info --- //
		r.Route("/info", func(r chi.Router) {
			r.Get("/", action.IndexInfo)
		})

		// --- Auth --- //
		r.Route("/auth", func(r chi.Router) {
			r.Post("/sign-up", action.PostUser)
			r.Post("/sign-in", action.PostSession)
			r.Post("/sign-out", action.DeleteSession)
		})

		// --- Budget --- //
		r.Route("/budgets", func(r chi.Router) {
			r.Use(action.BudgetAccountCTX)
			r.Route("/", func(r chi.Router) {
				r.Get("/", action.IndexBudgets)
				r.Get("/current", action.GetCurrentBudget)
				r.Post("/", action.PostBudget)
				r.Route("/{budgetUID}", func(r chi.Router) {
					r.Use(action.BudgetCTX)
					r.Get("/", action.GetBudget)

					// --- Income --- //
					r.Route("/income-list", func(r chi.Router) {
						r.Get("/", action.GetIncomeList)
						r.Post("/", action.PostIncome)
						r.Route("/{incomeID}", func(r chi.Router) {
							r.Use(action.IncomeCTX)
							r.Get("/", action.GetIncome)
							r.Patch("/", action.PatchIncome)
							r.Delete("/", action.DeleteIncome)
						})
					})

					// --- Expense --- //
					r.Route("/expenses", func(r chi.Router) {
						r.Get("/", action.IndexExpenses)
						r.Post("/", action.PostExpense)
						r.Route("/{expenseID}", func(r chi.Router) {
							r.Use(action.ExpenseCTX)
							r.Get("/", action.GetExpense)
							r.Patch("/", action.PatchExpense)
							r.Delete("/", action.DeleteExpense)
						})
					})
				})
			})
		})
	})

	return r
}
