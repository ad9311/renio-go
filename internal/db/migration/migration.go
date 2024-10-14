package migration

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pressly/goose/v3"
)

func Migrate() {
	db, err := goose.OpenDBWithDriver("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to open DB: %v\n", err)
	}
	defer db.Close()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to find current working directory: %v\n", err)
	}

	dir := filepath.Join(cwd, "./db/migrations")

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("failed to read migrations directory: %v\n", err)
	}

	if len(files) == 0 {
		fmt.Print("! Migrations directory is empty, skipping migrations\n\n")
		return
	}

	if err := goose.Up(db, dir); err != nil {
		log.Fatalf("failed to apply migrations: %v\n", err)
	}
}
