package svc

import "github.com/ad9311/renio-go/internal/model"

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
