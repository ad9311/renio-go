package migration

import (
	"database/sql"
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

func Up() error {
	var migExec MigExec
	if err := migExec.setUp(); err != nil {
		defer migExec.DB.Close()
		return err
	}

	if migExec.Skip {
		console.Info(emptyMigrations)
		return nil
	}

	if err := goose.Up(migExec.DB, migExec.MigDir); err != nil {
		return err
	}

	return nil
}

func Reset() error {
	var migExec MigExec
	if err := migExec.setUp(); err != nil {
		defer migExec.DB.Close()
		return err
	}

	if migExec.Skip {
		console.Info(emptyMigrations)
		return nil
	}

	if err := goose.Reset(migExec.DB, migExec.MigDir); err != nil {
		return err
	}

	return nil
}

// --- Helpers --- //

func (m *MigExec) setUp() error {
	db, err := goose.OpenDBWithDriver("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	migDir, err := dir.MigrationsDir()
	if err != nil {
		return err
	}

	files, err := os.ReadDir(migDir)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		m.Skip = true
	}

	m.DB = db
	m.MigDir = migDir

	return nil
}
