package svc

import (
	"github.com/ad9311/renio-go/internal/model"
)

type BudgetWithEntries struct {
	model.Budget
	IncomeList model.IncomeList
	Expenses   model.Expenses
}

func FindBudgets(budgetAccountID int) (model.Budgets, error) {
	budgets := model.Budgets{}
	if err := budgets.Index(budgetAccountID); err != nil {
		return budgets, err
	}

	return budgets, nil
}

func FindBudget(budget model.Budget) (BudgetWithEntries, error) {
	budgetWithEntries := BudgetWithEntries{
		Budget:     budget,
		IncomeList: model.IncomeList{},
		Expenses:   model.Expenses{},
	}

	if err := budget.Validate(); err != nil {
		return budgetWithEntries, err
	}

	if err := budgetWithEntries.IncomeList.Index(budget.ID); err != nil {
		return budgetWithEntries, err
	}

	if err := budgetWithEntries.Expenses.Index(budget.ID); err != nil {
		return budgetWithEntries, err
	}

	return budgetWithEntries, nil
}
