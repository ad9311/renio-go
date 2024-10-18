package model

import (
	"context"
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

type Incomes []Income

type IncomeFormData struct {
	Amount       float32 `json:"amount"`
	Description  string  `json:"description"`
	EntryClassID int     `json:"entryClassId"`
}

func (i *Income) Insert(budgetID int, entryClassID int) error {
	query := `INSERT INTO incomes (amount, description, budget_id, entry_class_id)
						VALUES ($1, $2, $3, $4) RETURNING *`

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
	query := `SELECT * FROM incomes WHERE id = $1`

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
						WHERE id = $4 RETURNING *`

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

func (i *Income) Delete() error {
	query := "DELETE FROM incomes WHERE id = $1 RETURNING *"

	if err := i.queryIncome(query, i.ID); err != nil {
		return err
	}

	return nil
}

// --- Helpers --- //

func (i *Income) queryIncome(query string, params ...any) error {
	pool := db.GetPool()
	ctx := context.Background()

	err := pool.QueryRow(ctx, query, params...).Scan(spreadIncomeValues(i)...)
	if err != nil {
		return err
	}

	return nil
}

func spreadIncomeValues(income *Income) []any {
	return []any{
		&income.ID,
		&income.Amount,
		&income.Description,
		&income.BudgetID,
		&income.EntryClassID,
		&income.CreatedAt,
		&income.UpdatedAt,
	}
}
