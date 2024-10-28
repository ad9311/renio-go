package action

import (
	"encoding/json"
	"net/http"

	"github.com/ad9311/renio-go/internal/eval"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/svc"
	"github.com/ad9311/renio-go/internal/vars"
)

// --- Actions --- //

func GetIncomeList(w http.ResponseWriter, r *http.Request) {
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
	errEval, ok := err.(*eval.ErrEval)
	if ok {
		WriteError(w, errEval.Issues, http.StatusBadRequest)
		return
	}
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
	budget := r.Context().Value(vars.BudgetKey).(model.Budget)

	var incomeFormData model.IncomeFormData
	if err := json.NewDecoder(r.Body).Decode(&incomeFormData); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	err := svc.UpdateIncome(&income, incomeFormData, budget)
	errEval, ok := err.(*eval.ErrEval)
	if ok {
		WriteError(w, errEval.Issues, http.StatusBadRequest)
		return
	}
	if err != nil {
		WriteError(w, ErrorToSlice(err), http.StatusBadRequest)
		return
	}

	WriteOK(w, income, http.StatusOK)
}

func DeleteIncome(w http.ResponseWriter, r *http.Request) {
	income := r.Context().Value(vars.IncomeKey).(model.Income)
	budget := r.Context().Value(vars.BudgetKey).(model.Budget)

	if err := svc.DeleteIncome(income, budget); err != nil {
		WriteError(w, ErrorToSlice(err), http.StatusBadRequest)
		return
	}

	WriteOK(w, income, http.StatusOK)
}
