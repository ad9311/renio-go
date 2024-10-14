package action

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/conf"
	"github.com/ad9311/renio-go/internal/model"
)

func IndexBudgets(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(conf.UserIDContext).(int)
	var budgetAccount model.BudgetAccount
	if err := budgetAccount.FindByUserID(userID); err != nil {
		WriteError(w, []string{"no budget account found"}, http.StatusInternalServerError)
		return
	}

	var budgets model.Budgets
	if err := budgets.Index(budgetAccount.ID); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusNotFound)
		return
	}

	WriteOK(w, budgets, http.StatusOK)
}

func PostBudget(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(conf.UserIDContext).(int)
	var budgetAccount model.BudgetAccount
	if err := budgetAccount.FindByUserID(userID); err != nil {
		WriteError(w, []string{"no budget account found"}, http.StatusInternalServerError)
		return
	}

	var budget model.Budget
	if err := budget.Create(budgetAccount.ID); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusNotFound)
		return
	}

	WriteOK(w, budget, http.StatusCreated)
}

func GetBudget(w http.ResponseWriter, r *http.Request) {
	WriteOK(w, "", http.StatusCreated)
}
