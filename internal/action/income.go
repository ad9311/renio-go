package action

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/eval"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/svc"
	"github.com/ad9311/renio-go/internal/vars"
)

// --- Actions --- //

func GetIncomeList(w http.ResponseWriter, r *http.Request) {
	budget := r.Context().Value(vars.BudgetKey).(model.Budget)

	errResponse := ErrorResponse{}
	incomeList := model.IncomeList{}
	if err := incomeList.Index(budget.ID); err != nil {
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}

	dataResponse := DataResponse{
		Content: Content{
			"incomeList": incomeList,
		},
	}

	WriteOK(w, dataResponse)
}

func PostIncome(w http.ResponseWriter, r *http.Request) {
	budget := r.Context().Value(vars.BudgetKey).(model.Budget)

	errResponse := ErrorResponse{}
	var incomeFormData model.IncomeFormData
	if err := DecodeJSON(r.Body, &incomeFormData); err != nil {
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}

	income, err := svc.CreateIncome(incomeFormData, budget)
	errEval, ok := err.(*eval.ErrEval)
	if ok {
		errResponse.AppendIssues(errEval.Issues)
		WriteError(w, errResponse)
		return
	}
	if err != nil {
		WriteError(w, errResponse)
		return
	}

	dataResponse := DataResponse{
		Content: Content{
			"income": income,
		},
		Status: http.StatusCreated,
	}

	WriteOK(w, dataResponse)
}

func GetIncome(w http.ResponseWriter, r *http.Request) {
	income := r.Context().Value(vars.IncomeKey).(model.Income)

	dataResponse := DataResponse{
		Content: Content{
			"income": income,
		},
	}

	WriteOK(w, dataResponse)
}

func PatchIncome(w http.ResponseWriter, r *http.Request) {
	income := r.Context().Value(vars.IncomeKey).(model.Income)
	budget := r.Context().Value(vars.BudgetKey).(model.Budget)

	errResponse := ErrorResponse{}
	var incomeFormData model.IncomeFormData
	if err := DecodeJSON(r.Body, &incomeFormData); err != nil {
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}

	err := svc.UpdateIncome(&income, incomeFormData, budget)
	errEval, ok := err.(*eval.ErrEval)
	if ok {
		errResponse.AppendIssues(errEval.Issues)
		WriteError(w, errResponse)
		return
	}
	if err != nil {
		errResponse.Append(err)
		return
	}

	dataResponse := DataResponse{
		Content: Content{
			"income": income,
		},
	}

	WriteOK(w, dataResponse)
}

func DeleteIncome(w http.ResponseWriter, r *http.Request) {
	income := r.Context().Value(vars.IncomeKey).(model.Income)
	budget := r.Context().Value(vars.BudgetKey).(model.Budget)

	errResponse := ErrorResponse{}
	if err := svc.DeleteIncome(income, budget); err != nil {
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}

	dataResponse := DataResponse{
		Content: Content{
			"income": income,
		},
	}

	WriteOK(w, dataResponse)
}
