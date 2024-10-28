package svc

import (
	"github.com/ad9311/renio-go/internal/model"
)

func UpdateIncome(
	income *model.Income,
	incomeFormData model.IncomeFormData,
	budget model.Budget,
) error {
	if err := incomeFormData.Validate(); err != nil {
		return err
	}

	prevAmount := income.Amount
	if err := income.Update(incomeFormData); err != nil {
		return err
	}

	if err := budget.UpdateOnEntry(income.Amount, prevAmount, 0); err != nil {
		return err
	}

	if err := budget.UpdateOnIncome(income.Amount, prevAmount, 0); err != nil {
		return err
	}

	return nil
}
