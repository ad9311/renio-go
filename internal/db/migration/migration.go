package migration

import (
	"fmt"
	"log"
	"os"

	"github.com/ad9311/renio-go/internal/dir"
	"github.com/pressly/goose/v3"
)

func Migrate(databaseURL string) {
	db, err := goose.OpenDBWithDriver("pgx", databaseURL)
	if err != nil {
		log.Fatalf("failed to open DB: %v\n", err)
	}
	defer db.Close()

	migDir, err := dir.MigrationsDir()
	if err != nil {
		log.Fatalf("failed to find migrations directory: %v\n", err)
	}

	files, err := os.ReadDir(migDir)
	if err != nil {
		log.Fatalf("failed to read migrations directory: %v\n", err)
	}

	if len(files) == 0 {
		fmt.Print("! Migrations directory is empty, skipping migrations\n\n")
		return
	}

	if err := goose.Up(db, migDir); err != nil {
		log.Fatalf("failed to apply migrations: %v\n", err)
	}
}
