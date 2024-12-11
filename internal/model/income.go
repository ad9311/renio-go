package model

import (
	"time"

	"github.com/ad9311/renio-go/internal/app"
)

type Income struct {
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

type IncomeList []Income

type IncomeFormData struct {
	Amount       float32
	Description  string
	EntryClassID int
}

// --- Query functions --- //

func (il *IncomeList) Index(budgetID int) error {
	query := `
	SELECT incomes.*,
	entry_classes.name AS entry_class_name,
	entry_classes.uid AS entry_class_uid
	FROM incomes
	INNER JOIN entry_classes ON incomes.entry_class_id = entry_classes.id
	WHERE budget_id = $1
	`

	var incomeList []any
	queryExec := app.QueryExe{
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

func (i *Income) Insert(incomeFormData IncomeFormData, budgetID int) error {
	query := `
	INSERT INTO incomes (amount, description, entry_class_id, budget_id)
	VALUES ($1, $2, $3, $4)
	RETURNING *
	`

	queryExec := app.QueryExe{
		QueryStr: query,
		QueryArgs: []any{
			incomeFormData.Amount,
			incomeFormData.Description,
			incomeFormData.EntryClassID,
			budgetID,
		},
		Model: Income{},
	}

	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := i.saveIncomeFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (i *Income) SelectByID(id int) error {
	query := `
	SELECT incomes.*,
	entry_classes.name AS entry_class_name,
	entry_classes.uid AS entry_class_uid
	FROM incomes
	INNER JOIN entry_classes ON incomes.entry_class_id = entry_classes.id
	WHERE incomes.id = $1
	`

	queryExec := app.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{id},
		Model:     Income{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := i.saveIncomeFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (i *Income) Update(incomeFormData IncomeFormData) error {
	query := `
	UPDATE incomes SET amount = $1, description = $2, entry_class_id = $3
	WHERE id = $4
	RETURNING *
	`

	queryExec := app.QueryExe{
		QueryStr: query,
		QueryArgs: []any{
			incomeFormData.Amount,
			incomeFormData.Description,
			incomeFormData.EntryClassID,
			i.ID,
		},
		Model: Income{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := i.saveIncomeFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (i *Income) Delete() error {
	query := "DELETE FROM incomes WHERE id = $1 RETURNING *"

	queryExec := app.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{i.ID},
		Model:     Income{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := i.saveIncomeFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (i *Income) FindLast(budgetID int) error {
	query := `
	SELECT incomes.*,
	entry_classes.name AS entry_class_name,
	entry_classes.uid AS entry_class_uid
	FROM incomes
	INNER JOIN entry_classes ON incomes.entry_class_id = entry_classes.id
	WHERE budget_id = $1
	ORDER BY id DESC
	LIMIT 1
	`

	queryExec := app.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{budgetID},
		Model:     Income{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := i.saveIncomeFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

// --- Helpers --- //

func (i *Income) saveIncomeFromDB(queryExec app.QueryExe) error {
	value, ok := queryExec.Model.(*Income)
	if !ok {
		return ErrIncompleteQuery{}
	}
	*i = *value

	return nil
}
