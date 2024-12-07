package handler

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/svc"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var budgetAccount model.BudgetAccount
	userID := getCurrentUserId(ctx)
	if err := budgetAccount.SelectByUserID(userID); err != nil {
		writeNotFound(w, ctx)
		return
	}

	budgetSummary, err := svc.FindBudgetSummary(budgetAccount.ID)
	if err != nil {
		writeNotFound(w, ctx)
		return
	}

	getAppData(ctx)["budget"] = budgetSummary
	writeTemplate(w, ctx, "home/index")
}
