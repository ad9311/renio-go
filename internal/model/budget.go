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

const budgetColumns = `id, uid, balance, total_income, total_expenses, transaction_count, income_count, expense_count`

// Query functions //

func (bs *Budgets) Index(budgetAccountID int) error {
	condition := "budget_account_id = $1"
	query := fmt.Sprintf("SELECT %s FROM budgets WHERE %s", budgetColumns, condition)

	if err := bs.queryBudgets(query, budgetAccountID); err != nil {
		return err
	}

	return nil
}

func (b *Budget) Insert(budgetAccountID int) error {
	query := `INSERT INTO budgets (uid, budget_account_id) VALUES ($1, $2) RETURNING`
	query = fmt.Sprintf("%s %s", query, budgetColumns)

	b.setCurrentUID(budgetAccountID)
	if err := b.queryBudget(query, b.UID, budgetAccountID); err != nil {
		return err
	}

	return nil
}

func (b *Budget) SelectByUID(budgetAccountID int, uid string) error {
	condition := "budget_account_id = $1 AND uid = $2"
	query := fmt.Sprintf("SELECT %s FROM budgets WHERE %s", budgetColumns, condition)
	if err := b.queryBudget(query, budgetAccountID, uid); err != nil {
		return err
	}

	return nil
}

func (b *Budget) SelectCurrent(budgetAccountID int) error {
	condition := "budget_account_id = $1 AND uid = $2"
	query := fmt.Sprintf("SELECT %s FROM budgets WHERE %s", budgetColumns, condition)
	b.setCurrentUID(budgetAccountID)
	if err := b.queryBudget(query, budgetAccountID, b.UID); err != nil {
		return err
	}

	return nil
}

// --- Helpers --- //

func (b *Budget) setCurrentUID(budgetAccountID int) {
	currentTime := time.Now()
	year := currentTime.Local().Year()
	month := currentTime.Local().Month()
	uid := fmt.Sprintf("%d-%d-%d", budgetAccountID, year, month)
	b.UID = uid
}

func (bs *Budgets) queryBudgets(query string, params ...any) error {
	pool := db.GetPool()
	ctx := context.Background()

	rows, err := pool.Query(ctx, query, params...)
	if err != nil {
		return err
	}
	defer rows.Close()

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

func (b *Budget) queryBudget(query string, params ...any) error {
	pool := db.GetPool()
	ctx := context.Background()

	err := pool.QueryRow(ctx, query, params...).Scan(
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
