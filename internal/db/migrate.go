package db

import (
	"os"
	"path/filepath"

	"github.com/ad9311/renio-go/internal/conf"
	"github.com/pressly/goose/v3"
)

func Migrate() error {
	db, err := goose.OpenDBWithDriver("pgx", conf.GetEnv().DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	dir := filepath.Join(cwd, "./db/migrations")

	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return nil
	}

	if err := goose.Up(db, dir); err != nil {
		return err
	}

	return nil
}
