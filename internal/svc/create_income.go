package svc

import "github.com/ad9311/renio-go/internal/model"

func CreateIncome(incomeFormData model.IncomeFormData, budget model.Budget) (model.Income, error) {
	// TODO: income validation

	var income model.Income
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
