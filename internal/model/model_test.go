package model_test

import (
	"log"
	"os"
	"testing"

	"github.com/ad9311/renio-go/internal/app"
	"github.com/ad9311/renio-go/internal/db/migration"
)

func TestMain(m *testing.M) {
	if err := app.Init(); err != nil {
		log.Fatal(err.Error())
	}

	code := m.Run()
	cleanUp()
	os.Exit(code)
}

func cleanUp() {
	if err := migration.Reset(); err != nil {
		panic(err.Error())
	}
}
