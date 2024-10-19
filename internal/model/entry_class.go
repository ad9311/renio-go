package model

import (
	"time"

	"github.com/ad9311/renio-go/internal/db"
)

type EntryClass struct {
	ID        int    `json:"id"`
	UID       string `json:"uid"`
	Name      string `json:"name"`
	Group     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EntryClasses []EntryClass

const (
	IncomeGroupName  = "income"
	ExpenseGroupName = "expense"
)

var EntryClassGroupNames = map[int]string{
	0: ExpenseGroupName,
	1: IncomeGroupName,
}

// --- Query --- //

func (e *EntryClass) Insert() error {
	query := `INSERT INTO entry_classes (uid, name, "group")
            VALUES ($1, $2, $3) RETURNING *`

	queryExec := db.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{e.UID, e.Name, e.Group},
		Model:     EntryClass{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	return nil
}

func (e *EntryClass) InsertIfNotExists() error {
	query := `INSERT INTO entry_classes (uid, name, "group")
						VALUES ($1, $2, $3) RETURNING *
						ON CONFLICT (uid) DO NOTHING`

	queryExec := db.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{e.UID, e.Name, e.Group},
		Model:     EntryClass{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	return nil
}

func (e *EntryClass) GroupName() string {
	return EntryClassGroupNames[e.Group]
}
