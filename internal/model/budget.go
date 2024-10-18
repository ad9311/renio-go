package model

import (
	"context"
	"fmt"
	"time"

	"github.com/ad9311/renio-go/internal/db"
)

type Budget struct {
	ID              int       `json:"id"`
	UID             string    `json:"uid"`
	Balance         float32   `json:"balance"`
	TotalIncome     float32   `json:"totalIncome"`
	TotalExpenses   float32   `json:"totalExpenses"`
	EntryCount      int       `json:"entryCount"`
	IncomeCount     int       `json:"incomeCount"`
	ExpenseCount    int       `json:"expenseCount"`
	BudgetAccountID int       `json:"budgetAccountId"`
	CreatedAt       time.Time `json:"createAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type Budgets []Budget

const budgetQuery1 = "SELECT * FROM budgets WHERE"

// Query functions //

func (bs *Budgets) Index(budgetAccountID int) error {
	query := fmt.Sprintf("%s budget_account_id = $1", budgetQuery1)

	if err := bs.queryBudgets(query, budgetAccountID); err != nil {
		return err
	}

	return nil
}

func (b *Budget) SelectByUID(uid string) error {
	query := fmt.Sprintf("%s uid = $1", budgetQuery1)
	if err := b.queryBudget(query, uid); err != nil {
		return err
	}

	return nil
}

func (b *Budget) SelectCurrent(budgetAccountID int) error {
	query := fmt.Sprintf("%s uid = $1", budgetQuery1)
	b.setCurrentUID(budgetAccountID)
	if err := b.queryBudget(query, b.UID); err != nil {
		return err
	}

	return nil
}

func (b *Budget) Insert(budgetAccountID int) error {
	query := "INSERT INTO budgets (uid, budget_account_id) VALUES ($1, $2) RETURNING *"

	b.setCurrentUID(budgetAccountID)
	if err := b.queryBudget(query, b.UID, budgetAccountID); err != nil {
		return err
	}

	return nil
}

func (b *Budget) OnIncomeInsert(incomeAmount float32) error {
	columns := "balance = $1, total_income = $2, entry_count = $3, income_count = $4"
	query := fmt.Sprintf("UPDATE budgets SET %s WHERE id = $5 RETURNING *", columns)

	b.setBalance(incomeAmount, 0)
	b.setTotalIncome(incomeAmount, 0)
	b.addToEntryCount(1)
	b.addToIncomeCount(1)

	if err := b.queryBudget(
		query,
		b.Balance,
		b.TotalIncome,
		b.EntryCount,
		b.IncomeCount,
		b.ID,
	); err != nil {
		return err
	}

	return nil
}

func (b *Budget) OnIncomeUpdate(prevIncomeAmount float32, incomeAmount float32) error {
	columns := "balance = $1, total_income = $2"
	query := fmt.Sprintf("UPDATE budgets SET %s WHERE id = $3 RETURNING *", columns)

	b.setBalance(incomeAmount, prevIncomeAmount)
	b.setTotalIncome(incomeAmount, prevIncomeAmount)

	if err := b.queryBudget(
		query,
		b.Balance,
		b.TotalIncome,
		b.ID,
	); err != nil {
		return err
	}

	return nil
}

func (b *Budget) OnIncomeDelete(incomeAmount float32) error {
	columns := "balance = $1, total_income = $2, entry_count = $3, income_count = $4"
	query := fmt.Sprintf("UPDATE budgets SET %s WHERE id = $5 RETURNING *", columns)

	b.setBalance(0, incomeAmount)
	b.setTotalIncome(0, incomeAmount)
	b.addToEntryCount(-1)
	b.addToIncomeCount(-1)

	if err := b.queryBudget(
		query,
		b.Balance,
		b.TotalIncome,
		b.EntryCount,
		b.IncomeCount,
		b.ID,
	); err != nil {
		return err
	}

	return nil
}

// --- Helpers --- //

func (b *Budget) setCurrentUID(budgetAccountID int) {
	currentTime := time.Now()
	year := currentTime.Local().Year()
	month := currentTime.Local().Month()
	uid := fmt.Sprintf("%d-%d-%02d", budgetAccountID, year, month)
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
		if err := rows.Scan(spreadBudgetValues(&budget)...); err != nil {
			return nil
		}

		*bs = append(*bs, budget)
	}

	return nil
}

func (b *Budget) queryBudget(query string, params ...any) error {
	pool := db.GetPool()
	ctx := context.Background()

	err := pool.QueryRow(ctx, query, params...).Scan(spreadBudgetValues(b)...)
	if err != nil {
		return err
	}

	return nil
}

func (b *Budget) setBalance(credit float32, debit float32) {
	b.Balance = b.Balance + (credit - debit)
}

func (b *Budget) setTotalIncome(credit float32, debit float32) {
	b.TotalIncome = b.TotalIncome + (credit - debit)
}

func (b *Budget) addToEntryCount(change int) {
	b.EntryCount = b.EntryCount + change
}

func (b *Budget) addToIncomeCount(change int) {
	b.IncomeCount = b.IncomeCount + change
}

func spreadBudgetValues(budget *Budget) []any {
	return []any{
		&budget.ID,
		&budget.UID,
		&budget.Balance,
		&budget.TotalIncome,
		&budget.TotalExpenses,
		&budget.EntryCount,
		&budget.IncomeCount,
		&budget.ExpenseCount,
		&budget.BudgetAccountID,
		&budget.CreatedAt,
		&budget.UpdatedAt,
	}
}
