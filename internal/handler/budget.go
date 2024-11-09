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
		budgetUID := chi.URLParam(r, "budgetUID")

		var budget model.Budget
		err := budget.SelectByUID(budgetUID)
		if err == pgx.ErrNoRows {
			return
		}
		if err != nil {
			return
		}

		ctx := context.WithValue(r.Context(), vars.BudgetKey, budget)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetBudgets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	budgetAccount := ctx.Value(vars.BudgetAccountKey).(model.BudgetAccount)

	budgets, err := svc.FindBudgets(budgetAccount.ID)
	if err != nil {
		writeTemplate(w, r, "error/index")
		return
	}

	GetAppData(ctx)["budgets"] = budgets
	writeTemplate(w, r, "budgets/index")
}
