package router

import (
	"context"
	"net/http"

	"github.com/ad9311/renio-go/internal/action"
	"github.com/ad9311/renio-go/internal/conf"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/go-chi/chi/v5"
)

func BudgetAccountCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(conf.UserIDContext).(int)

		var budgetAccount model.BudgetAccount
		if err := budgetAccount.SelectByUserID(userID); err != nil {
			action.WriteError(w, []string{"user not signed in"}, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), conf.BudgetAccountContext, budgetAccount)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func BudgetCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(conf.UserIDContext).(int)

		var budgetAccount model.BudgetAccount
		if err := budgetAccount.SelectByUserID(userID); err != nil {
			action.WriteError(w, []string{"user not signed in"}, http.StatusUnauthorized)
			return
		}

		budgetUID := chi.URLParam(r, "budgetUID")
		var budget model.Budget
		if err := budget.SelectByUID(budgetAccount.ID, budgetUID); err != nil {
			action.WriteError(w, []string{"budget not found"}, http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), conf.BudgetContext, budget)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
