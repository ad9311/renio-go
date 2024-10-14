package model

import (
	"context"
	"fmt"
	"time"

	"github.com/ad9311/renio-go/internal/db"
)

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

type Budgets []Budget

const budgetSelectedColumns = `id, uid, balance, total_income, total_expenses, transaction_count, income_count, expense_count`

// Query functions //

func (bs *Budgets) Index(budgetAccountID int) error {
	pool := db.GetPool()
	ctx := context.Background()
	query := `SELECT id, uid, balance, total_income, total_expenses,
						transaction_count, income_count, expense_count
						FROM budgets WHERE budget_account_id = $1`

	rows, err := pool.Query(ctx, query, budgetAccountID)
	if err != nil {
		return err
	}
	defer rows.Close()

	// var budgets []*Budget
	for rows.Next() {
		var budget Budget
		err := rows.Scan(
			&budget.ID,
			&budget.UID,
			&budget.Balance,
			&budget.TotalIncome,
			&budget.TotalExpenses,
			&budget.TransactionCount,
			&budget.IncomeCount,
			&budget.ExpenseCount,
		)
		if err != nil {
			return nil
		}

		*bs = append(*bs, budget)
	}

	return nil
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

func (b *Budget) FindByUID(uid string) error {
	pool := db.GetPool()
	ctx := context.Background()
	query := ``
	err := pool.QueryRow(ctx, query, uid).Scan(
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
