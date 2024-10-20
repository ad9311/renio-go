package action

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/conf"
	"github.com/ad9311/renio-go/internal/model"
)

// --- Actions --- //

func IndexBudgets(w http.ResponseWriter, r *http.Request) {
	budgetAccount := r.Context().Value(conf.BudgetAccountContext).(model.BudgetAccount)

	var budgets model.Budgets
	if err := budgets.Index(budgetAccount.ID); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusNotFound)
		return
	}

	WriteOK(w, budgets, http.StatusOK)
}

func PostBudget(w http.ResponseWriter, r *http.Request) {
	budgetAccount := r.Context().Value(conf.BudgetAccountContext).(model.BudgetAccount)

	var budget model.Budget
	if err := budget.Insert(budgetAccount.ID); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusNotFound)
		return
	}

	WriteOK(w, budget, http.StatusCreated)
}

func GetBudget(w http.ResponseWriter, r *http.Request) {
	budget := r.Context().Value(conf.BudgetContext).(model.Budget)
	WriteOK(w, budget, http.StatusOK)
}

func GetCurrentBudget(w http.ResponseWriter, r *http.Request) {
	budgetAccount := r.Context().Value(conf.BudgetAccountContext).(model.BudgetAccount)

	var budget model.Budget
	if err := budget.SelectCurrent(budgetAccount.ID); err != nil {
		WriteError(w, []string{""}, http.StatusNotFound)
	}

	WriteOK(w, budget, http.StatusOK)
}
