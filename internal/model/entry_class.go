package model

import (
	"context"

	"github.com/ad9311/renio-go/internal/db"
)

type EntryClass struct {
	ID        int    `json:"id"`
	UID       string `json:"uid"`
	Name      string `json:"name"`
	Group     int    `json:"group"`
	GroupName string `json:"groupName"`
}

type EntryClasses []EntryClass

const (
	IncomeGroupName  = "income"
	ExpenseGroupName = "expense"
)

var EntryClassGroupNames = map[int]string{
	0: IncomeGroupName,
	1: ExpenseGroupName,
}

// --- Query --- //

func (e *EntryClass) Insert() error {
	pool := db.GetPool()
	ctx := context.Background()
	query := `INSERT INTO entry_classes (uid, name) VALUES ($1, $2, $3)`

	if _, err := pool.Exec(ctx, query, e.UID, e.Name, e.Group); err != nil {
		return err
	}
	e.setGroupName()

	return nil
}

// --- Helpers --- //

func (e *EntryClass) setGroupName() {
	e.GroupName = EntryClassGroupNames[e.Group]
}
