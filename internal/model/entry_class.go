package model

import (
	"time"

	"github.com/ad9311/renio-go/internal/db"
)

type EntryClass struct {
	ID        int    `json:"id"`
	UID       string `json:"uid"`
	Name      string `json:"name"`
	Type      int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EntryClasses []EntryClass

const (
	IncomeTypeName  = "income"
	ExpenseTypeName = "expense"
)

var EntryClassTypeNames = map[int]string{
	0: ExpenseTypeName,
	1: IncomeTypeName,
}

// --- Query --- //

func (e *EntryClass) Insert() error {
	query := `INSERT INTO entry_classes (uid, name, type) VALUES ($1, $2, $3) RETURNING *`

	queryExec := db.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{e.UID, e.Name, e.Type},
		Model:     EntryClass{},
	}
	if err := queryExec.QueryRow(); err != nil {
		return err
	}

	if err := e.saveEntryClassFromDB(queryExec); err != nil {
		return err
	}

	return nil
}

func (e *EntryClass) InsertIfNotExists() error {
	query := `INSERT INTO entry_classes (uid, name, type) VALUES ($1, $2, $3) ON CONFLICT (uid) DO NOTHING`

	queryExec := db.QueryExe{
		QueryStr:  query,
		QueryArgs: []any{e.UID, e.Name, e.Type},
	}
	if err := queryExec.Exec(); err != nil {
		return err
	}

	return nil
}

func (e *EntryClass) TypeName() string {
	return EntryClassTypeNames[e.Type]
}

// --- Helpers  --- //

func (e *EntryClass) saveEntryClassFromDB(queryExec db.QueryExe) error {
	value, ok := queryExec.Model.(*EntryClass)
	if !ok {
		return ErrIncompleteQuery{}
	}
	*e = *value

	return nil
}
