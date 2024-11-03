package handler

import (
	"context"
	"net/http"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/vars"
	"github.com/jackc/pgx/v5"
)

func BudgetAccountCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var budgetAccount model.BudgetAccount
		err := budgetAccount.SelectByUserID(0)
		if err == pgx.ErrNoRows {
			return
		}
		if err != nil {
			return
		}

		ctx := context.WithValue(r.Context(), vars.BudgetAccountKey, budgetAccount)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
