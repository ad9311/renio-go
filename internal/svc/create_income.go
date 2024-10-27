package svc

import (
	"github.com/ad9311/renio-go/internal/eval"
	"github.com/ad9311/renio-go/internal/model"
)

type IncomeData struct {
	Income model.Income
	Issues eval.Issues
}

func CreateIncome(incomeFormData model.IncomeFormData, budget model.Budget) (IncomeData, error) {
	var createdIncome IncomeData

	issues := incomeFormData.Validate()
	if issues != nil {
		createdIncome.Issues = issues
		return createdIncome, nil
	}
	var income model.Income
	if err := income.Insert(incomeFormData, budget.ID); err != nil {
		return createdIncome, err
	}

	if err := budget.UpdateOnEntry(income.Amount, 0, 1); err != nil {
		return createdIncome, err
	}

	if err := budget.UpdateOnIncome(income.Amount, 0, 1); err != nil {
		return createdIncome, err
	}

	createdIncome.Income = income

	return createdIncome, nil
}
