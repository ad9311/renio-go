package main

import (
	"os"

	"github.com/ad9311/renio-go/internal/console"
	"github.com/ad9311/renio-go/internal/db"
	"github.com/ad9311/renio-go/internal/db/migration"
	"github.com/ad9311/renio-go/internal/db/seed"
	"github.com/ad9311/renio-go/internal/envs"
)

func main() {
	console.AppName()

	envs.Init()

	migrate := os.Getenv("MIGRATE")
	seeds := os.Getenv("SEED")

	db.Init()

	if migrate == "on" {
		migration.Up()
	}

	if seeds == "on" {
		seed.RunSeeds()
	}

	Serve()
}
