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
	if err := godotenv.Load("../../.env.test"); err != nil {
		console.Error(err.Error())
		os.Exit(1)
	}

	console.ResetEnv()

	db.Init()
	migration.Up()
	seed.Run()

	code := m.Run()

	migration.Reset()

	os.Exit(code)
}
