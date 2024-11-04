package handler

import (
	"context"
	"net/http"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/vars"
)

func BudgetAccountCTX(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var budgetAccount model.BudgetAccount
		userID := GetCurrentUserId(r.Context())
		err := budgetAccount.SelectByUserID(userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeTemplate(w, r, "error/index")
			return
		}

		ctx := context.WithValue(r.Context(), vars.BudgetAccountKey, budgetAccount)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
