package model_test

import (
	"os"
	"testing"

	"github.com/ad9311/renio-go/internal/console"
	"github.com/ad9311/renio-go/internal/db"
	"github.com/ad9311/renio-go/internal/db/migration"
	"github.com/ad9311/renio-go/internal/db/seed"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../.env"); err != nil {
		console.Error(err.Error())
		os.Exit(1)
	}

	console.ResetEnv()

	databaseURL := os.Getenv("TEST_DATABASE_URL")
	db.Init(databaseURL)
	migration.Up(databaseURL)
	seed.RunSeeds()

	code := m.Run()

	migration.Reset(databaseURL)

	os.Exit(code)
}
