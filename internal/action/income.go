package action

import (
	"encoding/json"
	"net/http"

	"github.com/ad9311/renio-go/internal/conf"
	"github.com/ad9311/renio-go/internal/model"
)

// --- Actions --- //

func IndexIncomeList(w http.ResponseWriter, r *http.Request) {
	budget := r.Context().Value(conf.BudgetContext).(model.Budget)

	var incomeList model.IncomeList
	if err := incomeList.Index(budget.ID); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return

	}

	WriteOK(w, incomeList, http.StatusCreated)
}

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
		WriteError(w, []string{err.Error()}, http.StatusInternalServerError)
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

	prevIncomeAmount := income.Amount
	if err := income.Update(incomeFormData); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	if err := budget.OnIncomeUpdate(prevIncomeAmount, income.Amount); err != nil {
		WriteError(w, []string{"failed to updated budget"}, http.StatusInternalServerError)
		return
	}

	WriteOK(w, income, http.StatusOK)
}

func DeleteIncome(w http.ResponseWriter, r *http.Request) {
	income := r.Context().Value(conf.IncomeContext).(model.Income)
	budget := r.Context().Value(conf.BudgetContext).(model.Budget)

	if err := income.Delete(); err != nil {
		WriteError(w, []string{"failed to delete income"}, http.StatusInternalServerError)
		return
	}

	if err := budget.OnIncomeDelete(income.Amount); err != nil {
		WriteError(w, []string{"failed to update budget"}, http.StatusInternalServerError)
		return
	}

	WriteOK(w, income, http.StatusOK)
}
