package model_test

import "testing"

func TestBudgetInsert(t *testing.T) {
	budget := PrepareBudget(t)

	if budget.Balance != 0 {
		t.Errorf("expected budget with balanace %f, got %f", 0.0, budget.Balance)
	}

	if budget.TotalIncome != 0 {
		t.Errorf("expected budget with total income %f, got %f", 0.0, budget.TotalIncome)
	}

	if budget.TotalExpenses != 0 {
		t.Errorf("expected budget with total expenses %f, got %f", 0.0, budget.TotalExpenses)
	}

	if budget.EntryCount != 0 {
		t.Errorf("expected budget with entry count %d, got %d", 0, budget.EntryCount)
	}

	if budget.IncomeCount != 0 {
		t.Errorf("expected budget with income count %d, got %d", 0, budget.IncomeCount)
	}

	if budget.ExpenseCount != 0 {
		t.Errorf("expected budget with expense count %d, got %d", 0, budget.ExpenseCount)
	}
}
