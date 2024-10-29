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

	var expenses model.Expenses
	if err := expenses.Index(budget.ID); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return

	}

	WriteOK(w, expenses, http.StatusCreated)
}

func PostExpense(w http.ResponseWriter, r *http.Request) {
	budget := r.Context().Value(vars.BudgetKey).(model.Budget)

	var expenseFormData model.ExpenseFormData
	if err := json.NewDecoder(r.Body).Decode(&expenseFormData); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	expense, err := svc.CreateExpense(expenseFormData, budget)
	errEval, ok := err.(*eval.ErrEval)
	if ok {
		WriteError(w, errEval.Issues, http.StatusBadRequest)
		return
	}
	if err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	WriteOK(w, expense, http.StatusCreated)
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
	expense := r.Context().Value(vars.ExpenseKey).(model.Expense)
	WriteOK(w, expense, http.StatusOK)
}

func PatchExpense(w http.ResponseWriter, r *http.Request) {
	expense := r.Context().Value(vars.ExpenseKey).(model.Expense)
	// budget := r.Context().Value(vars.BudgetKey).(model.Budget)

	var expenseFormData model.ExpenseFormData
	if err := json.NewDecoder(r.Body).Decode(&expenseFormData); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	// prevExpenseAmount := expense.Amount
	// if err := expense.Update(expenseFormData); err != nil {
	// 	WriteError(w, []string{err.Error()}, http.StatusBadRequest)
	// 	return
	// }

	// if err := budget.OnExpenseUpdate(prevExpenseAmount, expense.Amount); err != nil {
	// 	WriteError(w, []string{"failed to updated budget"}, http.StatusInternalServerError)
	// 	return
	// }

	WriteOK(w, expense, http.StatusOK)
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	expense := r.Context().Value(vars.ExpenseKey).(model.Expense)
	// budget := r.Context().Value(vars.BudgetKey).(model.Budget)

	if err := expense.Delete(); err != nil {
		WriteError(w, []string{"failed to delete expense"}, http.StatusInternalServerError)
		return
	}

	// if err := budget.OnExpenseDelete(expense.Amount); err != nil {
	// 	WriteError(w, []string{"failed to update budget"}, http.StatusInternalServerError)
	// 	return
	// }

	WriteOK(w, expense, http.StatusOK)
}
