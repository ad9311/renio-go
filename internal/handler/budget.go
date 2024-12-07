package handler

import (
	"context"
	"net/http"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/svc"
	"github.com/ad9311/renio-go/internal/vars"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func BudgetCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		budgetAccount := ctx.Value(vars.BudgetAccountKey).(model.BudgetAccount)
		budgetUID := chi.URLParam(r, "budgetUID")

		var budget model.Budget
		err := budget.SelectByUID(budgetUID, budgetAccount.ID)
		if err == pgx.ErrNoRows {
			writeNotFound(w, ctx)
			return
		}
		if err != nil {
			errStr := []string{err.Error()}
			writeInternalError(w, ctx, errStr)
			return
		}

		ctx = context.WithValue(ctx, vars.BudgetKey, budget)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetBudgets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budgetAccount := ctx.Value(vars.BudgetAccountKey).(model.BudgetAccount)

	budgets, err := svc.FindBudgets(budgetAccount.ID)
	if err != nil {
		errStr := []string{err.Error()}
		writeInternalError(w, ctx, errStr)
		return
	}

	getAppData(ctx)["budgets"] = budgets
	writeTemplate(w, ctx, "budgets/index")
}

func GetBudget(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budget := ctx.Value(vars.BudgetKey).(model.Budget)

	budgetWithEntries, err := svc.FindBudget(budget)
	if err == pgx.ErrNoRows {
		writeNotFound(w, ctx)
		return
	}
	if err != nil {
		errStr := []string{err.Error()}
		writeInternalError(w, ctx, errStr)
		return
	}

	getAppData(ctx)["budget"] = budgetWithEntries
	writeTemplate(w, ctx, "budgets/show")
}
