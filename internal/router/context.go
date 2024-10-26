package router

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ad9311/renio-go/internal/action"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/vars"
	"github.com/go-chi/chi/v5"
)

func BudgetAccountCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedJWT := r.Context().Value(vars.AllowedJWTKey).(model.AllowedJWT)

		var budgetAccount model.BudgetAccount
		if err := budgetAccount.SelectByUserID(allowedJWT.UserID); err != nil {
			action.WriteError(w, []string{"user not signed in"}, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), vars.BudgetAccountKey, budgetAccount)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func BudgetCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		budgetUID := chi.URLParam(r, "budgetUID")

		var budget model.Budget
		if err := budget.SelectByUID(budgetUID); err != nil {
			action.WriteError(w, []string{err.Error()}, http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), vars.BudgetKey, budget)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func IncomeCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		incomeID := chi.URLParam(r, "incomeID")
		id, _ := strconv.Atoi(incomeID)
		income := model.Income{
			ID: id,
		}
		if err := income.SelectByID(); err != nil {
			action.WriteError(w, []string{"income not found"}, http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), vars.IncomeKey, income)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ExpenseCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expenseID := chi.URLParam(r, "expenseID")
		id, _ := strconv.Atoi(expenseID)
		expense := model.Expense{
			ID: id,
		}
		if err := expense.SelectByID(); err != nil {
			action.WriteError(w, []string{"expense not found"}, http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), vars.ExpenseKey, expense)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
