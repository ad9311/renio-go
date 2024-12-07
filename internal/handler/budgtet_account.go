package handler

import (
	"context"
	"net/http"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/vars"
)

func BudgetAccountCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var budgetAccount model.BudgetAccount
		userID := getCurrentUserId(ctx)

		if err := budgetAccount.SelectByUserID(userID); err != nil {
			writeInternalError(w, ctx, []string{err.Error()})
			return
		}

		ctx = context.WithValue(ctx, vars.BudgetAccountKey, budgetAccount)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
