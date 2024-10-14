package ctrl

import (
	"context"
	"net/http"

	"github.com/ad9311/renio-go/internal/conf"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/go-chi/chi/v5"
)

func BudgetRouter(r chi.Router) func(r chi.Router) {
	return func(r chi.Router) {
		r.Use(BudgetCTX)
		r.Get("/", indexBudgets) // r.Route("/{budgetUID}", func(r chi.Router) {
		// 	r.Use(BudgetCTX)
		// 	r.Get("/", findBudget)
		// })
		r.Post("/", createBudget)
	}
}

func BudgetCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		budgetUID := chi.URLParam(r, "budgetUID")
		userID := r.Context().Value(conf.UserIDContext).(int)

		var budgetAccount model.BudgetAccount
		if err := budgetAccount.FindByUserID(userID); err != nil {
			WriteError(w, []string{"budget not found"}, http.StatusNotFound)
			return
		}

		var budget model.Budget
		if err := budget.FindByUID(budgetAccount.UserID, budgetUID); err != nil {
			WriteError(w, []string{"budget not found"}, http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), BudgetUIDParam, budget)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// actions //

func indexBudgets(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(conf.UserIDContext).(int)
	var budgetAccount model.BudgetAccount
	if err := budgetAccount.FindByUserID(userID); err != nil {
		WriteError(w, []string{"no budget account found"}, http.StatusInternalServerError)
		return
	}

	var budgets model.Budgets
	if err := budgets.Index(budgetAccount.ID); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusNotFound)
		return
	}

	WriteOK(w, budgets, http.StatusOK)
}

func createBudget(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(conf.UserIDContext).(int)
	var budgetAccount model.BudgetAccount
	if err := budgetAccount.FindByUserID(userID); err != nil {
		WriteError(w, []string{"no budget account found"}, http.StatusInternalServerError)
		return
	}

	var budget model.Budget
	if err := budget.Create(budgetAccount.ID); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusNotFound)
		return
	}

	WriteOK(w, budget, http.StatusCreated)
}

func findBudget(w http.ResponseWriter, r *http.Request) {
	budget := r.Context().Value(BudgetUIDParam).(*model.Budget)
	WriteOK(w, budget, http.StatusCreated)
}
