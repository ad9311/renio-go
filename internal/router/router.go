package router

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/action"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RoutesHandler() http.Handler {
	r := chi.NewRouter()

	// --- Middlewares --- //
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
			r.Post("/sign-in", action.PostSession)
			r.Post("/sign-up", action.PostUser)
		})

		// --- Budget --- //
		r.Route("/budgets", func(r chi.Router) {
			r.Use(BudgetAccountCTX)
			r.Route("/", func(r chi.Router) {
				r.Get("/", action.IndexBudgets)
				r.Get("/current", action.GetCurrentBudget)
				r.Post("/", action.PostBudget)
				r.Route("/{budgetUID}", func(r chi.Router) {
					r.Use(BudgetCTX)
					r.Get("/", action.GetBudget)

					// --- Income ---
					r.Route("/income-list", func(r chi.Router) {
						r.Post("/", action.PostIncome)
						r.Route("/{incomeID}", func(r chi.Router) {
							r.Use(IncomeCTX)
							r.Get("/", action.GetIncome)
							r.Patch("/", action.PatchIncome)
						})
					})
				})
			})
		})
	})

	return r
}
