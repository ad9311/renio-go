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

// Query functions //

func (bs *Budgets) Index(budgetAccountID int) error {
	query := "SELECT * FROM budgets WHERE budget_account_id = $1"

	var budgets []any
	queryExec := db.QueryExe{
		QueryStr:   query,
		QueryArgs:  []any{budgetAccountID},
		Model:      Budget{},
		ModelSlice: &budgets,
	}

	if err := queryExec.Query(); err != nil {
		return err
	}

	fmt.Println(&queryExec.ModelSlice)

	for _, b := range budgets {
		budget := b.(*Budget) // Type assertion
		*bs = append(*bs, *budget)
	}

	// if err := bs.queryBudgets(query, budgetAccountID); err != nil {
	//	return err
	// }

	return nil
}

func (b *Budget) SelectByUID(uid string) error {
	dbExec := db.QueryExe{
		QueryStr:  "SELECT * FROM budgets WHERE uid = $1",
		QueryArgs: []any{uid},
		ScanArgs:  spreadBudgetValues(b),
	}
	if err := db.QueryRow(dbExec); err != nil {
		return err
	}

	return nil
}

func (b *Budget) SelectCurrent(budgetAccountID int) error {
	b.setCurrentUID(budgetAccountID)
	dbExec := db.QueryExe{
		QueryStr:  "SELECT * FROM budgets WHERE uid = $1",
		QueryArgs: []any{b.UID},
		ScanArgs:  spreadBudgetValues(b),
	}
	if err := db.QueryRow(dbExec); err != nil {
		return err
	}

	return nil
}

func (b *Budget) Insert(budgetAccountID int) error {
	query := "INSERT INTO budgets (uid, budget_account_id) VALUES ($1, $2) RETURNING *"
	dbExec := db.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{b.UID, budgetAccountID},
		ScanArgs:  spreadBudgetValues(b),
	}
	if err := db.QueryRow(dbExec); err != nil {
		return err
	}

	return nil
}

func (b *Budget) OnIncomeInsert(incomeAmount float32) error {
	query := `UPDATE budgets SET
            balance = $1, total_income = $2, entry_count = $3, income_count = $4
            WHERE ID = $5 RETURNING *`

	b.setBalance(incomeAmount, 0)
	b.setTotalIncome(incomeAmount, 0)
	b.addToEntryCount(1)
	b.addToIncomeCount(1)

	dbExec := db.QueryExe{
		QueryStr: query,
		QueryArgs: []any{
			b.Balance,
			b.TotalIncome,
			b.EntryCount,
			b.IncomeCount,
			b.ID,
		},
		ScanArgs: spreadBudgetValues(b),
	}
	if err := db.QueryRow(dbExec); err != nil {
		return err
	}

	return nil
}

func (b *Budget) OnIncomeUpdate(prevIncomeAmount float32, incomeAmount float32) error {
	query := `UPDATE budgets SET balance = $1, total_income = $2 WHERE id = $3 RETURNING *`

	b.setBalance(incomeAmount, prevIncomeAmount)
	b.setTotalIncome(incomeAmount, prevIncomeAmount)
	dbExec := db.QueryExe{
		QueryStr: query,
		QueryArgs: []any{
			b.Balance,
			b.TotalIncome,
			b.ID,
		},
		ScanArgs: spreadBudgetValues(b),
	}
	if err := db.QueryRow(dbExec); err != nil {
		return err
	}

	return nil
}

func (b *Budget) OnIncomeDelete(incomeAmount float32) error {
	query := `UPDATE budgets SET
            balance = $1, total_income = $2, entry_count = $3, income_count = $4
            WHERE ID = $5 RETURNING *`

	b.setBalance(0, incomeAmount)
	b.setTotalIncome(0, incomeAmount)
	b.addToEntryCount(-1)
	b.addToIncomeCount(-1)

	dbExec := db.QueryExe{
		QueryStr: query,
		QueryArgs: []any{
			b.Balance,
			b.TotalIncome,
			b.EntryCount,
			b.IncomeCount,
			b.ID,
		},
		ScanArgs: spreadBudgetValues(b),
	}
	if err := db.QueryRow(dbExec); err != nil {
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
