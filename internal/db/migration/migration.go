package migration

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ad9311/renio-go/internal/console"
	"github.com/ad9311/renio-go/internal/dir"
	"github.com/pressly/goose/v3"
)

type MigExec struct {
	DB     *sql.DB
	MigDir string
	Skip   bool
}

const emptyMigrations = "Migrations directory is empty, skipping migrations"

func Up() {
	var migExec MigExec
	migExec.setUp()
	defer migExec.DB.Close()

	if migExec.Skip {
		console.Info(emptyMigrations)
		return
	}

	if err := goose.Up(migExec.DB, migExec.MigDir); err != nil {
		console.Fatal(fmt.Sprintf("could not migrate up: %s", err.Error()))
	}
}

func Reset() {
	var migExec MigExec
	migExec.setUp()
	defer migExec.DB.Close()

	if migExec.Skip {
		console.Info(emptyMigrations)
		return
	}

	if err := goose.Reset(migExec.DB, migExec.MigDir); err != nil {
		console.Fatal(fmt.Sprintf("could not reset database: %s", err.Error()))
	}
}

// --- Helpers --- //

func (m *MigExec) setUp() {
	db, err := goose.OpenDBWithDriver("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		console.Fatal(fmt.Sprintf("failed to open database, %s", err.Error()))
	}

	migDir, err := dir.MigrationsDir()
	if err != nil {
		console.Fatal(fmt.Sprintf("failed to find migrations directory: %s", err.Error()))
	}

	files, err := os.ReadDir(migDir)
	if err != nil {
		console.Fatal(fmt.Sprintf("failed to read migrations directory: %s", err.Error()))
	}

	if len(files) == 0 {
		m.Skip = true
	}

	m.DB = db
	m.MigDir = migDir
}
