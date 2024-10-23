package action

import (
	"encoding/json"
	"net/http"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/vars"
)

// --- Actions --- //

func IndexExpenses(w http.ResponseWriter, r *http.Request) {
	budget := r.Context().Value(vars.BudgetContext).(model.Budget)

	var expenses model.Expenses
	if err := expenses.Index(budget.ID); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return

	}

	WriteOK(w, expenses, http.StatusCreated)
}

func PostExpense(w http.ResponseWriter, r *http.Request) {
	budget := r.Context().Value(vars.BudgetContext).(model.Budget)

	var expenseFormData model.ExpenseFormData
	if err := json.NewDecoder(r.Body).Decode(&expenseFormData); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	expense := model.Expense{
		Amount:      expenseFormData.Amount,
		Description: expenseFormData.Description,
	}
	if err := expense.Insert(budget.ID, expenseFormData.EntryClassID); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	if err := budget.OnExpenseInsert(expense.Amount); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusInternalServerError)
		return
	}

	WriteOK(w, expense, http.StatusCreated)
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
	expense := r.Context().Value(vars.ExpenseContext).(model.Expense)
	WriteOK(w, expense, http.StatusOK)
}

func PatchExpense(w http.ResponseWriter, r *http.Request) {
	expense := r.Context().Value(vars.ExpenseContext).(model.Expense)
	budget := r.Context().Value(vars.BudgetContext).(model.Budget)

	var expenseFormData model.ExpenseFormData
	if err := json.NewDecoder(r.Body).Decode(&expenseFormData); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	prevExpenseAmount := expense.Amount
	if err := expense.Update(expenseFormData); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	if err := budget.OnExpenseUpdate(prevExpenseAmount, expense.Amount); err != nil {
		WriteError(w, []string{"failed to updated budget"}, http.StatusInternalServerError)
		return
	}

	WriteOK(w, expense, http.StatusOK)
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	expense := r.Context().Value(vars.ExpenseContext).(model.Expense)
	budget := r.Context().Value(vars.BudgetContext).(model.Budget)

	if err := expense.Delete(); err != nil {
		WriteError(w, []string{"failed to delete expense"}, http.StatusInternalServerError)
		return
	}

	if err := budget.OnExpenseDelete(expense.Amount); err != nil {
		WriteError(w, []string{"failed to update budget"}, http.StatusInternalServerError)
		return
	}

	WriteOK(w, expense, http.StatusOK)
}
