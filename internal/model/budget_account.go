package model

import (
	"context"

	"github.com/ad9311/renio-go/internal/db"
)

type BudgetAccount struct {
	ID     int
	UserID int
}

func (ba *BudgetAccount) Insert(userID int) error {
	pool := db.GetPool()
	ctx := context.Background()
	query := `INSERT INTO budget_accounts (user_id)
						VALUES ($1)
						RETURNING id, user_id`

	err := pool.QueryRow(ctx, query, userID).Scan(&ba.ID, &ba.UserID)
	if err != nil {
		return err
	}

	return nil
}
