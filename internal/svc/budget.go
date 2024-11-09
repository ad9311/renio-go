package svc

import (
	"github.com/ad9311/renio-go/internal/model"
	"github.com/jackc/pgx/v5"
)

type BudgetWithEntries struct {
	model.Budget
	IncomeList model.IncomeList
	Expenses   model.Expenses
}

type BudgetSummary struct {
	model.Budget
	LastIncome  model.Income
	LastExpense model.Expense
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

func FindBudgetSummary(budgetAccountID int) (BudgetSummary, error) {
	var budgetSumarry BudgetSummary
	var budget model.Budget

	err := budget.SelectCurrent(budgetAccountID)
	if err == pgx.ErrNoRows {
		budget, err = CreateCurrentBudget(budgetAccountID)
		if err != nil {
			return budgetSumarry, err
		}
	}
	if err != nil && err != pgx.ErrNoRows {
		return budgetSumarry, err
	}

	var lastIncome model.Income
	err = lastIncome.FindLast(budget.ID)
	if err != nil && err != pgx.ErrNoRows {
		return budgetSumarry, err
	}

	var lastExpense model.Expense
	err = lastExpense.FindLast(budget.ID)
	if err != nil && err != pgx.ErrNoRows {
		return budgetSumarry, err
	}

	budgetSumarry.Budget = budget
	budgetSumarry.LastIncome = lastIncome
	budgetSumarry.LastExpense = lastExpense

	return budgetSumarry, nil
}

func CreateCurrentBudget(budgetAccountID int) (model.Budget, error) {
	var budget model.Budget

	if err := budget.Insert(budgetAccountID); err != nil {
		return budget, err
	}

	return budget, nil
}
