package migration

import (
	"fmt"
	"log"
	"os"

	"github.com/ad9311/renio-go/internal/console"
	"github.com/ad9311/renio-go/internal/dir"
	"github.com/pressly/goose/v3"
)

func Migrate(databaseURL string) {
	db, err := goose.OpenDBWithDriver("pgx", databaseURL)
	if err != nil {
		console.Fatal(fmt.Sprintf("failed to open database, %s", err.Error()))
	}
	defer db.Close()

	migDir, err := dir.MigrationsDir()
	if err != nil {
		console.Fatal(fmt.Sprintf("failed to find migrations directory: %s", err.Error()))
	}

	files, err := os.ReadDir(migDir)
	if err != nil {
		console.Fatal(fmt.Sprintf("failed to read migrations directory: %s", err.Error()))
	}

	if len(files) == 0 {
		console.Info("Migrations directory is empty, skipping migrations")
		return
	}

	if err := goose.Up(db, migDir); err != nil {
		log.Fatalf("failed to apply migrations: %v\n", err)
		console.Fatal(fmt.Sprintf("failed to apply migrations: %s", err.Error()))
	}
}
