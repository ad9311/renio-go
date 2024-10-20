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

func main() {
	console.AppName()

	env := os.Getenv("ENV")

	if env != "production" {
		if err := godotenv.Load(); err != nil {
			console.Fatal(fmt.Sprintf("could not load .env file, %s", err.Error()))
		}
		env = os.Getenv("ENV")
		console.Success(fmt.Sprintf("Loaded .env file in %s environment", env))
	}

	databaseURL := os.Getenv("DATABASE_URL")
	migrate := os.Getenv("MIGRATE")
	seeds := os.Getenv("SEED")

	db.Init(databaseURL)

	if migrate == "on" {
		migration.Up(databaseURL)
	}

	if seeds == "on" {
		seed.RunSeeds()
	}

	Serve()
}
