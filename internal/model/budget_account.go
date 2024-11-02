package model

import (
	"time"

	"github.com/ad9311/renio-go/internal/db"
)

type BudgetAccount struct {
	ID        int
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// --- Query Functions --- //

func (b *BudgetAccount) Insert(userID int) error {
	query := "INSERT INTO budget_accounts (user_id) VALUES ($1) RETURNING *"

	queryExec := db.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{userID},
		Model:     BudgetAccount{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := b.saveBudgetAccountFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (b *BudgetAccount) SelectByUserID(userID int) error {
	query := "SELECT * FROM budget_accounts WHERE user_id = $1"
	queryExec := db.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{userID},
		Model:     BudgetAccount{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := b.saveBudgetAccountFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

// --- Helpers --- //

func (b *BudgetAccount) saveBudgetAccountFromDB(queryExec db.QueryExe) error {
	value, ok := queryExec.Model.(*BudgetAccount)
	if !ok {
		return ErrIncompleteQuery{}
	}
	*b = *value

	return nil
}
