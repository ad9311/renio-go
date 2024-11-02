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

func IndexExpenses(w http.ResponseWriter, r *http.Request) {
	budget := r.Context().Value(vars.BudgetKey).(model.Budget)

	errResponse := ErrorResponse{}
	expenses := model.Expenses{}
	if err := expenses.Index(budget.ID); err != nil {
		errResponse.Append(err)
		WriteError(w, errResponse)
		return

	}

	dataResponse := DataResponse{
		Content: Content{
			"expenses": expenses,
		},
	}

	WriteOK(w, dataResponse)
}

func PostExpense(w http.ResponseWriter, r *http.Request) {
	budget := r.Context().Value(vars.BudgetKey).(model.Budget)

	errResponse := ErrorResponse{}
	var expenseFormData model.ExpenseFormData
	if err := json.NewDecoder(r.Body).Decode(&expenseFormData); err != nil {
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}

	expense, err := svc.CreateExpense(expenseFormData, budget)
	errEval, ok := err.(*eval.ErrEval)
	if ok {
		errResponse.AppendIssues(errEval.Issues)
		WriteError(w, errResponse)
		return
	}
	if err != nil {
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}

	dataResponse := DataResponse{
		Content: Content{
			"expense": expense,
		},
	}

	WriteOK(w, dataResponse)
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
	expense := r.Context().Value(vars.ExpenseKey).(model.Expense)

	dataResponse := DataResponse{
		Content: Content{
			"expense": expense,
		},
	}

	WriteOK(w, dataResponse)
}

func PatchExpense(w http.ResponseWriter, r *http.Request) {
	expense := r.Context().Value(vars.ExpenseKey).(model.Expense)
	budget := r.Context().Value(vars.BudgetKey).(model.Budget)

	errResponse := ErrorResponse{}
	var expenseFormData model.ExpenseFormData
	if err := json.NewDecoder(r.Body).Decode(&expenseFormData); err != nil {
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}

	err := svc.UpdateExpense(&expense, expenseFormData, budget)
	errEval, ok := err.(*eval.ErrEval)
	if ok {
		errResponse.AppendIssues(errEval.Issues)
		WriteError(w, errResponse)
		return
	}
	if err != nil {
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}

	dataResponse := DataResponse{
		Content: Content{
			"expense": expense,
		},
	}

	WriteOK(w, dataResponse)
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	expense := r.Context().Value(vars.ExpenseKey).(model.Expense)
	budget := r.Context().Value(vars.BudgetKey).(model.Budget)

	errResponse := ErrorResponse{}
	if err := svc.DeleteExpense(expense, budget); err != nil {
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}

	dataResponse := DataResponse{
		Content: Content{
			"expense": expense,
		},
	}

	WriteOK(w, dataResponse)
}
