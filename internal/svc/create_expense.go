package svc

import "github.com/ad9311/renio-go/internal/model"

func CreateExpense(
	expenseFormData model.ExpenseFormData,
	budget model.Budget,
) (model.Expense, error) {
	var expense model.Expense

	if err := expenseFormData.Validate(); err != nil {
		return expense, err
	}

	if err := expense.Insert(expenseFormData, budget.ID); err != nil {
		return expense, err
	}

	if err := budget.UpdateOnEntry(0, expense.Amount, 1); err != nil {
		return expense, err
	}

	if err := budget.UpdateOnExpense(expense.Amount, 0, 1); err != nil {
		return expense, err
	}

	return expense, nil
}
