package model

import (
	"context"
	"fmt"

	"github.com/ad9311/renio-go/internal/db"
)

type Income struct {
	ID          int     `json:"id"`
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`
}

type Incomes []Income

type IncomeFormData struct {
	Amount       float32 `json:"amount"`
	Description  string  `json:"description"`
	EntryClassID int     `json:"entryClassId"`
}

const incomeColumns = `id, amount, description`

func (i *Income) Insert(budgetID int, entryClassID int) error {
	query := `INSERT INTO incomes (amount, description, budget_id, entry_class_id)
						VALUES ($1, $2, $3, $4) RETURNING`
	query = fmt.Sprintf("%s %s", query, incomeColumns)

	if err := i.queryIncome(
		query,
		i.Amount,
		i.Description,
		budgetID,
		entryClassID,
	); err != nil {
		return err
	}

	return nil
}

func (i *Income) SelectByID() error {
	query := `SELECT id, amount, description FROM incomes WHERE id = $1`

	if err := i.queryIncome(query, i.ID); err != nil {
		return err
	}

	return nil
}

func (i *Income) Update(incomeFormData IncomeFormData) error {
	query := `UPDATE incomes SET
						amount = $1,
						description = $2,
						entry_class_id = $3
						WHERE id = $4 RETURNING`
	query = fmt.Sprintf("%s %s", query, incomeColumns)

	if err := i.queryIncome(
		query,
		incomeFormData.Amount,
		incomeFormData.Description,
		incomeFormData.EntryClassID,
		i.ID,
	); err != nil {
		return err
	}

	return nil
}

// --- Helpers --- //

func (i *Income) queryIncome(query string, params ...any) error {
	pool := db.GetPool()
	ctx := context.Background()

	err := pool.QueryRow(ctx, query, params...).Scan(
		&i.ID,
		&i.Amount,
		&i.Description,
	)
	if err != nil {
		return err
	}

	return nil
}
