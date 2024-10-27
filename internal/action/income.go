package action

import (
	"encoding/json"
	"net/http"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/svc"
	"github.com/ad9311/renio-go/internal/vars"
)

// --- Actions --- //

func IndexIncomeList(w http.ResponseWriter, r *http.Request) {
	budget := r.Context().Value(vars.BudgetKey).(model.Budget)

	incomeList := model.IncomeList{}
	if err := incomeList.Index(budget.ID); err != nil {
		WriteError(w, ErrorToSlice(err), http.StatusBadRequest)
		return
	}

	WriteOK(w, incomeList, http.StatusOK)
}

func PostIncome(w http.ResponseWriter, r *http.Request) {
	budget := r.Context().Value(vars.BudgetKey).(model.Budget)

	var incomeFormData model.IncomeFormData
	if err := DecodeJSON(r.Body, &incomeFormData); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	income, err := svc.CreateIncome(incomeFormData, budget)
	if err != nil {
		WriteError(w, ErrorToSlice(err), http.StatusBadRequest)
		return
	}

	WriteOK(w, income, http.StatusCreated)
}

func GetIncome(w http.ResponseWriter, r *http.Request) {
	income := r.Context().Value(vars.IncomeKey).(model.Income)
	WriteOK(w, income, http.StatusOK)
}

func PatchIncome(w http.ResponseWriter, r *http.Request) {
	income := r.Context().Value(vars.IncomeKey).(model.Income)
	// budget := r.Context().Value(vars.BudgetKey).(model.Budget)

	var incomeFormData model.IncomeFormData
	if err := json.NewDecoder(r.Body).Decode(&incomeFormData); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	// prevIncomeAmount := income.Amount
	// if err := income.Update(incomeFormData); err != nil {
	// 	WriteError(w, []string{err.Error()}, http.StatusBadRequest)
	// 	return
	// }

	// if err := budget.OnIncomeUpdate(prevIncomeAmount, income.Amount); err != nil {
	// 	WriteError(w, []string{"failed to updated budget"}, http.StatusInternalServerError)
	// 	return
	// }

	WriteOK(w, income, http.StatusOK)
}

func DeleteIncome(w http.ResponseWriter, r *http.Request) {
	income := r.Context().Value(vars.IncomeKey).(model.Income)
	// budget := r.Context().Value(vars.BudgetKey).(model.Budget)

	if err := income.Delete(); err != nil {
		WriteError(w, []string{"failed to delete income"}, http.StatusInternalServerError)
		return
	}

	// if err := budget.OnIncomeDelete(income.Amount); err != nil {
	// 	WriteError(w, []string{"failed to update budget"}, http.StatusInternalServerError)
	// 	return
	// }

	WriteOK(w, income, http.StatusOK)
}
