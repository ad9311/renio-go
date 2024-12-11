package model

import (
	"time"

	"github.com/ad9311/renio-go/internal/app"
)

type Expense struct {
	ID             int
	Amount         float32
	Description    string
	BudgetID       int
	EntryClassID   int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	EntryClassName string
	EntryClassUID  string
}

type Expenses []Expense

type ExpenseFormData struct {
	Amount       float32
	Description  string
	EntryClassID int
}

// --- Query Functions -- //

func (es *Expenses) Index(budgetID int) error {
	query := `
	SELECT expenses.*,
	entry_classes.name AS entry_class_name,
	entry_classes.uid AS entry_class_uid
	FROM expenses
	INNER JOIN entry_classes ON expenses.entry_class_id = entry_classes.id
	WHERE budget_id = $1
	ORDER BY created_at DESC
	`

	var expenses []any
	queryExec := app.QueryExe{
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
	query := `
	INSERT INTO expenses (amount, description, entry_class_id, budget_id)
	VALUES ($1, $2, $3, $4)
	RETURNING *
	`

	queryExec := app.QueryExe{
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
	query := `
	SELECT expenses.*,
	entry_classes.name AS entry_class_name,
	entry_classes.uid AS entry_class_uid
	FROM expenses
	INNER JOIN entry_classes ON expenses.entry_class_id = entry_classes.id
	WHERE expenses.id = $1
	`

	queryExec := app.QueryExe{
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
	query := `
	UPDATE expenses SET amount = $1, description = $2, entry_class_id = $3
	WHERE id = $4
	RETURNING *
	`

	queryExec := app.QueryExe{
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

	queryExec := app.QueryExe{
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

func (e *Expense) FindLast(budgetID int) error {
	query := `
	SELECT expenses.*,
	entry_classes.name AS entry_class_name,
	entry_classes.uid AS entry_class_uid
	FROM expenses
	INNER JOIN entry_classes ON expenses.entry_class_id = entry_classes.id
	WHERE budget_id = $1
	ORDER BY id DESC
	LIMIT 1
	`

	queryExec := app.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{budgetID},
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

// --- Helpers --- //

func (e *Expense) saveExpenseFromDB(queryExec app.QueryExe) error {
	value, ok := queryExec.Model.(*Expense)
	if !ok {
		return ErrIncompleteQuery{}
	}
	*e = *value

	return nil
}
