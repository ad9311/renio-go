package ct

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/conf"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/go-chi/chi/v5"
)

func BudgetRouter(r chi.Router) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", createBudget)
	}
}

func createBudget(w http.ResponseWriter, r *http.Request) {
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
