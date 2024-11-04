package model

import (
	"fmt"
	"time"

	"github.com/ad9311/renio-go/internal/db"
	"github.com/ad9311/renio-go/internal/eval"
)

type Budget struct {
	ID              int
	UID             string
	Balance         float32
	TotalIncome     float32
	TotalExpenses   float32
	EntryCount      int
	IncomeCount     int
	ExpenseCount    int
	BudgetAccountID int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Budgets []Budget

// Query functions //

func (bs *Budgets) Index(budgetAccountID int) error {
	query := "SELECT * FROM budgets WHERE budget_account_id = $1 ORDER BY uid DESC"

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

	for _, b := range budgets {
		budget := b.(*Budget)
		*bs = append(*bs, *budget)
	}

	return nil
}

func (b *Budget) SelectByUID(uid string) error {
	queryExec := db.QueryExe{
		QueryStr:  "SELECT * FROM budgets WHERE uid = $1",
		QueryArgs: []any{uid},
		Model:     Budget{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := b.saveBudgetFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (b *Budget) SelectCurrent(budgetAccountID int) error {
	b.setCurrentUID(budgetAccountID)

	queryExec := db.QueryExe{
		QueryStr:  "SELECT * FROM budgets WHERE uid = $1",
		QueryArgs: []any{b.UID},
		Model:     Budget{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := b.saveBudgetFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (b *Budget) Insert(budgetAccountID int) error {
	query := "INSERT INTO budgets (uid, budget_account_id) VALUES ($1, $2) RETURNING *"
	b.setCurrentUID(budgetAccountID)
	queryExec := db.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{b.UID, budgetAccountID},
		Model:     Budget{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := b.saveBudgetFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (b *Budget) UpdateOnEntry(credit float32, debit float32, count int) error {
	query := "UPDATE budgets SET balance = $1, entry_count = $2 WHERE ID = $3 RETURNING *"

	b.setBalance(credit, debit)
	b.addToEntryCount(count)

	queryExec := db.QueryExe{
		QueryStr: query,
		QueryArgs: []any{
			b.Balance,
			b.EntryCount,
			b.ID,
		},
		Model: Budget{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := b.saveBudgetFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (b *Budget) UpdateOnIncome(credit float32, debit float32, count int) error {
	query := "UPDATE budgets SET total_income = $1, income_count = $2 WHERE ID = $3 RETURNING *"

	b.setTotalIncome(credit, debit)
	b.addToIncomeCount(count)

	queryExec := db.QueryExe{
		QueryStr: query,
		QueryArgs: []any{
			b.TotalIncome,
			b.IncomeCount,
			b.ID,
		},
		Model: Budget{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := b.saveBudgetFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (b *Budget) UpdateOnExpense(credit float32, debit float32, count int) error {
	query := "UPDATE budgets SET total_expenses = $1, expense_count = $2 WHERE ID = $3 RETURNING *"

	b.setTotalExpenses(credit, debit)
	b.addToExpenseCount(count)

	queryExec := db.QueryExe{
		QueryStr: query,
		QueryArgs: []any{
			b.TotalExpenses,
			b.IncomeCount,
			b.ID,
		},
		Model: Budget{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := b.saveBudgetFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

// --- Validations --- //

func (b *Budget) Validate() error {
	data := eval.ModelEval{
		Strings: []eval.String{
			{
				Name:    "Budget UID",
				Value:   b.UID,
				Pattern: `^\d-\d{4}-\d{2}$`,
			},
		},
		Ints: []eval.Int{
			{
				Name:     "Budget ID",
				Value:    b.ID,
				Positive: true,
			},
		},
	}
	if err := data.Validate(); err != nil {
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

func (b *Budget) setBalance(credit float32, debit float32) {
	b.Balance = b.Balance + (credit - debit)
}

func (b *Budget) setTotalIncome(credit float32, debit float32) {
	b.TotalIncome = b.TotalIncome + (credit - debit)
}

func (b *Budget) setTotalExpenses(credit float32, debit float32) {
	b.TotalExpenses = b.TotalExpenses + (credit - debit)
}

func (b *Budget) addToEntryCount(change int) {
	b.EntryCount = b.EntryCount + change
}

func (b *Budget) addToIncomeCount(change int) {
	b.IncomeCount = b.IncomeCount + change
}

func (b *Budget) addToExpenseCount(change int) {
	b.ExpenseCount = b.ExpenseCount + change
}

func (b *Budget) saveBudgetFromDB(queryExec db.QueryExe) error {
	value, ok := queryExec.Model.(*Budget)
	if !ok {
		return ErrIncompleteQuery{}
	}
	*b = *value

	return nil
}
