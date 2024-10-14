package model

import (
	"context"
	"fmt"
	"time"

	"github.com/ad9311/renio-go/internal/db"
)

// Query functions //

type Budget struct {
	ID               int     `json:"id"`
	UID              string  `json:"uid"`
	Balance          float64 `json:"balance"`
	TotalIncome      float64 `json:"totalIncome"`
	TotalExpenses    float64 `json:"totalExpenses"`
	TransactionCount int     `json:"transactionCount"`
	IncomeCount      int     `json:"incomeCount"`
	ExpenseCount     int     `json:"expenseCount"`
}

func (b *Budget) Create(budgetAccountID int) error {
	pool := db.GetPool()
	ctx := context.Background()
	query := `INSERT INTO budgets (uid, budget_account_id)
						VALUES ($1, $2)
						RETURNING id, uid, balance, total_income, total_expenses,
						transaction_count, income_count, expense_count`

	b.genUID(budgetAccountID)
	err := pool.QueryRow(ctx, query, b.UID, budgetAccountID).Scan(
		&b.ID,
		&b.UID,
		&b.Balance,
		&b.TotalIncome,
		&b.TotalExpenses,
		&b.TransactionCount,
		&b.IncomeCount,
		&b.ExpenseCount,
	)
	if err != nil {
		return err
	}

	return nil
}

// Helpers //

func (b *Budget) genUID(budgetAccountID int) {
	currentTime := time.Now()
	year := currentTime.Local().Year()
	month := currentTime.Local().Month()
	uid := fmt.Sprintf("%d-%d-%d", budgetAccountID, year, month)
	b.UID = uid
}
