package model

import (
	"time"

	"github.com/ad9311/renio-go/internal/db"
)

type Income struct {
	ID           int       `json:"id"`
	Amount       float32   `json:"amount"`
	Description  string    `json:"description"`
	BudgetID     int       `json:"budgetId"`
	EntryClassID int       `json:"entryClassId"`
	CreatedAt    time.Time `json:"createAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type IncomeList []Income

type IncomeFormData struct {
	Amount       float32 `json:"amount"`
	Description  string  `json:"description"`
	EntryClassID int     `json:"entryClassId"`
}

// --- Query functions --- //

func (il *IncomeList) Index(budgetID int) error {
	query := "SELECT * FROM incomes WHERE budget_id = $1 ORDER BY created_at DESC"

	var incomeList []any
	queryExec := db.QueryExe{
		QueryStr:   query,
		QueryArgs:  []any{budgetID},
		Model:      Income{},
		ModelSlice: &incomeList,
	}

	if err := queryExec.Query(); err != nil {
		return err
	}

	for _, i := range incomeList {
		income := i.(*Income)
		*il = append(*il, *income)
	}

	return nil
}

func (i *Income) Insert(budgetID int, entryClassID int) error {
	query := "INSERT INTO incomes (amount, description, budget_id, entry_class_id) VALUES ($1, $2, $3, $4) RETURNING *"

	queryExec := db.QueryExe{
		QueryStr: query,
		QueryArgs: []any{
			i.Amount,
			i.Description,
			budgetID,
			entryClassID,
		},
		Model: Income{},
	}

	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	value, ok := queryExec.Model.(*Income)
	if !ok {
		return ErrIncompleteQuery{}
	}
	*i = *value

	return nil
}

func (i *Income) SelectByID() error {
	query := "SELECT * FROM incomes WHERE id = $1"

	queryExec := db.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{i.ID},
		Model:     Income{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	value, ok := queryExec.Model.(*Income)
	if !ok {
		return ErrIncompleteQuery{}
	}
	*i = *value

	return nil
}

func (i *Income) Update(incomeFormData IncomeFormData) error {
	query := "UPDATE incomes SET amount = $1, description = $2, entry_class_id = $3 WHERE id = $4 RETURNING *"

	queryExec := db.QueryExe{
		QueryStr: query,
		QueryArgs: []any{
			incomeFormData.Amount,
			incomeFormData.Description,
			incomeFormData.EntryClassID,
		},
		Model: Income{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	value, ok := queryExec.Model.(*Income)
	if !ok {
		return ErrIncompleteQuery{}
	}
	*i = *value

	return nil
}

func (i *Income) Delete() error {
	query := "DELETE FROM incomes WHERE id = $1 RETURNING *"

	queryExec := db.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{i.ID},
		Model:     Income{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	value, ok := queryExec.Model.(*Income)
	if !ok {
		return ErrIncompleteQuery{}
	}
	*i = *value

	return nil
}
