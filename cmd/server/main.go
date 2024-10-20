package main

import (
	"fmt"
	"os"

	"github.com/ad9311/renio-go/internal/console"
	"github.com/ad9311/renio-go/internal/db"
	"github.com/ad9311/renio-go/internal/db/migration"
	"github.com/ad9311/renio-go/internal/db/seed"
	"github.com/joho/godotenv"
)

var (
	env         = os.Getenv("RENIO_ENV")
	databaseURL = os.Getenv("DATABASE_URL")
	migrate     = os.Getenv("MIGRATE")
	seeds       = os.Getenv("SEED")
)

func main() {
	console.AppName()

	if env != "production" {
		if err := godotenv.Load(); err != nil {
			console.Fatal(fmt.Sprintf("could not load .env file, %s", err.Error()))
		}
		env = os.Getenv("RENIO_ENV")
		console.Success(fmt.Sprintf("Loaded .env file in %s environment", env))
	}

	db.Init(databaseURL)

	if migrate == "on" {
		migration.Migrate(databaseURL)
	}

	if seeds == "on" {
		seed.RunSeeds()
	}

	Serve()
}
