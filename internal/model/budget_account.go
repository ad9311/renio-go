package model

import (
	"context"

	"github.com/ad9311/renio-go/internal/db"
)

type BudgetAccount struct {
	ID     int
	UserID int
}

// --- Query --- //

func (b *BudgetAccount) Insert(userID int) error {
	pool := db.GetPool()
	ctx := context.Background()
	query := `INSERT INTO budget_accounts (user_id)
						VALUES ($1)
						RETURNING id, user_id`

	err := pool.QueryRow(ctx, query, userID).Scan(&b.ID, &b.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (b *BudgetAccount) FindByUserID(userID int) error {
	query := `SELECT id, user_id FROM budget_accounts WHERE user_id = $1`
	pool := db.GetPool()
	ctx := context.Background()

	err := pool.QueryRow(ctx, query, userID).Scan(&b.ID, &b.UserID)
	if err != nil {
		return err
	}

	return nil
}
