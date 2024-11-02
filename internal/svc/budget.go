package svc

import (
	"github.com/ad9311/renio-go/internal/model"
)

type BudgetWithEntries struct {
	model.Budget
	IncomeList model.IncomeList
	Expenses   model.Expenses
}

func FindBudgetWithEntries(
	budget model.Budget,
	queries map[string]bool,
) (BudgetWithEntries, error) {
	budgetWithEntries := BudgetWithEntries{
		IncomeList: model.IncomeList{},
		Expenses:   model.Expenses{},
	}
	budgetWithEntries.Budget = budget

	if queries["income-list"] {
		if err := budgetWithEntries.IncomeList.Index(budget.ID); err != nil {
			return budgetWithEntries, err
		}
	}

	if queries["expenses"] {
		if err := budgetWithEntries.Expenses.Index(budget.ID); err != nil {
			return budgetWithEntries, err
		}
	}

	return budgetWithEntries, nil
}

func ParseBudgetEntriesQuery(query string) map[string]bool {
	options := map[string]bool{}

	if query == "income-list" {
		options["income-list"] = true
	}

	if query == "expenses" {
		options["expenses"] = true
	}

	if query == "full" {
		options["income-list"] = true
		options["expenses"] = true
	}

	return options
}
