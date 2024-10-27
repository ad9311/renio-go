package svc

import (
	"github.com/ad9311/renio-go/internal/eval"
	"github.com/ad9311/renio-go/internal/model"
)

func UpdateIncome(
	income *model.Income,
	incomeFormData model.IncomeFormData,
	budget model.Budget,
) (eval.Issues, error) {
	issues := incomeFormData.Validate()
	if issues != nil {
		return issues, nil
	}

	prevAmount := income.Amount
	if err := income.Update(incomeFormData); err != nil {
		return issues, err
	}

	if err := budget.UpdateOnEntry(income.Amount, prevAmount, 0); err != nil {
		return issues, err
	}

	if err := budget.UpdateOnIncome(income.Amount, prevAmount, 0); err != nil {
		return issues, err
	}

	return issues, nil
}
