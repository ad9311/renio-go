package handler

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/svc"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var budgetAccount model.BudgetAccount
	userID := GetCurrentUserId(ctx)
	if err := budgetAccount.SelectByUserID(userID); err != nil {
		GetAppData(ctx)["errors"] = []string{"could not select budget account for user"}
		w.WriteHeader(http.StatusInternalServerError)
		writeTemplate(w, r, "error/index")
		return
	}

	budgetSummary, err := svc.FindBudgetSummary(budgetAccount.ID)
	if err != nil {
		GetAppData(ctx)["errors"] = []string{err.Error()}
		writeTemplate(w, r, "error/index")
		return
	}

	GetAppData(ctx)["budget"] = budgetSummary
	writeTemplate(w, r, "home/index")
}
