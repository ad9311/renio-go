package model_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ad9311/renio-go/internal/db"
	"github.com/ad9311/renio-go/internal/db/migration"
	"github.com/ad9311/renio-go/internal/db/seed"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	databaseURL := os.Getenv("TEST_DATABASE_URL")
	db.Init(databaseURL)
	migration.Migrate(databaseURL)
	seed.RunSeeds()

	code := m.Run()
	os.Exit(code)
}
