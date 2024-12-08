package svc

import (
	"github.com/ad9311/renio-go/internal/model"
)

func CreateIncome(incomeFormData model.IncomeFormData, budget model.Budget) (model.Income, error) {
	var income model.Income

	if err := incomeFormData.Validate(); err != nil {
		return income, err
	}

	if err := income.Insert(incomeFormData, budget.ID); err != nil {
		return income, err
	}

	if err := budget.UpdateOnEntry(income.Amount, 0, 1); err != nil {
		return income, err
	}

	if err := budget.UpdateOnIncome(income.Amount, 0, 1); err != nil {
		return income, err
	}

	return income, nil
}

func UpdateIncome(
	income model.Income,
	incomeFormData model.IncomeFormData,
	budget model.Budget,
) (model.Income, error) {
	if err := incomeFormData.Validate(); err != nil {
		return income, err
	}

	prevAmount := income.Amount
	if err := income.Update(incomeFormData); err != nil {
		return income, err
	}

	if err := budget.UpdateOnEntry(income.Amount, prevAmount, 0); err != nil {
		return income, err
	}

	if err := budget.UpdateOnIncome(income.Amount, prevAmount, 0); err != nil {
		return income, err
	}

	return income, nil
}

func DeleteIncome(income model.Income, budget model.Budget) error {
	if err := income.Delete(); err != nil {
		return err
	}

	if err := budget.UpdateOnEntry(0, income.Amount, -1); err != nil {
		return err
	}

	if err := budget.UpdateOnIncome(0, income.Amount, -1); err != nil {
		return err
	}

	return nil
}
