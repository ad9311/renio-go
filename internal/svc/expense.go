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

func UpdateExpense(
	expense model.Expense,
	expenseformData model.ExpenseFormData,
	budget model.Budget,
) (model.Expense, error) {
	if err := expenseformData.Validate(); err != nil {
		return expense, err
	}

	prevAmount := expense.Amount
	if err := expense.Update(expenseformData); err != nil {
		return expense, err
	}

	if err := budget.UpdateOnEntry(prevAmount, expense.Amount, 0); err != nil {
		return expense, err
	}

	if err := budget.UpdateOnExpense(expense.Amount, prevAmount, 0); err != nil {
		return expense, err
	}

	return expense, nil
}

func DeleteExpense(expense model.Expense, budget model.Budget) error {
	if err := expense.Delete(); err != nil {
		return err
	}

	if err := budget.UpdateOnEntry(expense.Amount, 0, -1); err != nil {
		return err
	}

	if err := budget.UpdateOnExpense(0, expense.Amount, -1); err != nil {
		return err
	}

	return nil
}
