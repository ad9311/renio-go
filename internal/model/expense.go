package model

import (
	"time"

	"github.com/ad9311/renio-go/internal/db"
)

type Expense struct {
	ID           int       `json:"id"`
	Amount       float32   `json:"amount"`
	Description  string    `json:"description"`
	BudgetID     int       `json:"budgetId"`
	EntryClassID int       `json:"entryClassId"`
	CreatedAt    time.Time `json:"createAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Expenses []Expense

type ExpenseFormData struct {
	Amount       float32 `json:"amount"`
	Description  string  `json:"description"`
	EntryClassID int     `json:"entryClassId"`
}

// --- Query Functions -- //

func (es *Expenses) Index(budgetID int) error {
	query := "SELECT * FROM expenses WHERE budget_id = $1 ORDER BY created_at DESC"

	var expenses []any
	queryExec := db.QueryExe{
		QueryStr:   query,
		QueryArgs:  []any{budgetID},
		Model:      Expense{},
		ModelSlice: &expenses,
	}

	if err := queryExec.Query(); err != nil {
		return err
	}

	for _, e := range expenses {
		expense := e.(*Expense)
		*es = append(*es, *expense)
	}

	return nil
}

func (e *Expense) Insert(expenseFormData ExpenseFormData, budgetID int) error {
	query := "INSERT INTO expenses (amount, description, entry_class_id, budget_id) VALUES ($1, $2, $3, $4) RETURNING *"

	queryExec := db.QueryExe{
		QueryStr: query,
		QueryArgs: []any{
			expenseFormData.Amount,
			expenseFormData.Description,
			expenseFormData.EntryClassID,
			budgetID,
		},
		Model: Expense{},
	}

	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := e.saveExpenseFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (e *Expense) SelectByID(id int) error {
	query := "SELECT * FROM expenses WHERE id = $1"

	queryExec := db.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{id},
		Model:     Expense{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := e.saveExpenseFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (e *Expense) Update(expenseFormData ExpenseFormData) error {
	query := "UPDATE expenses SET amount = $1, description = $2, entry_class_id = $3 WHERE id = $4 RETURNING *"

	queryExec := db.QueryExe{
		QueryStr: query,
		QueryArgs: []any{
			expenseFormData.Amount,
			expenseFormData.Description,
			expenseFormData.EntryClassID,
			e.ID,
		},
		Model: Expense{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := e.saveExpenseFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (e *Expense) Delete() error {
	query := "DELETE FROM expenses WHERE id = $1 RETURNING *"

	queryExec := db.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{e.ID},
		Model:     Expense{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := e.saveExpenseFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (e *Expense) saveExpenseFromDB(queryExec db.QueryExe) error {
	value, ok := queryExec.Model.(*Expense)
	if !ok {
		return ErrIncompleteQuery{}
	}
	*e = *value

	return nil
}
