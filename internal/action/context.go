package action

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/vars"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func BudgetAccountCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedJWT := r.Context().Value(vars.AllowedJWTKey).(model.AllowedJWT)

		errResponse := ErrorResponse{}
		var budgetAccount model.BudgetAccount
		err := budgetAccount.SelectByUserID(allowedJWT.UserID)
		if err == pgx.ErrNoRows {
			err = fmt.Errorf("budget account not found")
			errResponse.Append(err)
			WriteError(w, errResponse)
			return
		}
		if err != nil {
			errResponse.Append(err)
			WriteError(w, errResponse)
			return
		}

		ctx := context.WithValue(r.Context(), vars.BudgetAccountKey, budgetAccount)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func BudgetCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		budgetUID := chi.URLParam(r, "budgetUID")

		errResponse := ErrorResponse{}
		var budget model.Budget
		err := budget.SelectByUID(budgetUID)
		if err == pgx.ErrNoRows {
			err = fmt.Errorf("budget not found")
			errResponse.Append(err)
			WriteError(w, errResponse)
			return
		}
		if err != nil {
			errResponse.Append(err)
			WriteError(w, errResponse)
			return
		}

		ctx := context.WithValue(r.Context(), vars.BudgetKey, budget)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func IncomeCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		incomeID := chi.URLParam(r, "incomeID")

		errResponse := ErrorResponse{}
		id, _ := strconv.Atoi(incomeID)
		var income model.Income

		err := income.SelectByID(id)
		if err == pgx.ErrNoRows {
			err = fmt.Errorf("income not found")
			errResponse.Append(err)
			WriteError(w, errResponse)
			return
		}
		if err != nil {
			errResponse.Append(err)
			WriteError(w, errResponse)
			return
		}

		ctx := context.WithValue(r.Context(), vars.IncomeKey, income)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ExpenseCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expenseID := chi.URLParam(r, "expenseID")

		errResponse := ErrorResponse{}
		id, _ := strconv.Atoi(expenseID)
		var expense model.Expense

		err := expense.SelectByID(id)
		if err == pgx.ErrNoRows {
			err = fmt.Errorf("expense not found")
			errResponse.Append(err)
			WriteError(w, errResponse)
			return
		}
		if err != nil {
			errResponse.Append(err)
			WriteError(w, errResponse)
			return
		}

		ctx := context.WithValue(r.Context(), vars.ExpenseKey, expense)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
