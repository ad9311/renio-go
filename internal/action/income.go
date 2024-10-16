package action

import (
	"encoding/json"
	"net/http"

	"github.com/ad9311/renio-go/internal/conf"
	"github.com/ad9311/renio-go/internal/model"
)

// --- Actions --- //

func PostIncome(w http.ResponseWriter, r *http.Request) {
	budget := r.Context().Value(conf.BudgetContext).(model.Budget)

	var incomeFormData model.IncomeFormData
	if err := json.NewDecoder(r.Body).Decode(&incomeFormData); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	income := model.Income{
		Amount:      incomeFormData.Amount,
		Description: incomeFormData.Description,
	}
	if err := income.Insert(budget.ID, incomeFormData.EntryClassID); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}
	if err := budget.OnIncomeInsert(income.Amount); err != nil {
		WriteError(w, []string{"failed to updated budget"}, http.StatusInternalServerError)
		return
	}

	WriteOK(w, income, http.StatusCreated)
}

func GetIncome(w http.ResponseWriter, r *http.Request) {
	income := r.Context().Value(conf.IncomeContext).(model.Income)
	WriteOK(w, income, http.StatusOK)
}

func PatchIncome(w http.ResponseWriter, r *http.Request) {
	income := r.Context().Value(conf.IncomeContext).(model.Income)
	budget := r.Context().Value(conf.BudgetContext).(model.Budget)

	var incomeFormData model.IncomeFormData
	if err := json.NewDecoder(r.Body).Decode(&incomeFormData); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	amount := income.Amount
	if err := income.Update(incomeFormData); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}
	if err := budget.OnIncomeUpdate(amount, income.Amount); err != nil {
		WriteError(w, []string{"failed to updated budget"}, http.StatusInternalServerError)
		return
	}

	WriteOK(w, income, http.StatusOK)
}
