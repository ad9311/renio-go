package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ad9311/renio-go/internal/db"
	"github.com/ad9311/renio-go/internal/db/migration"
	"github.com/ad9311/renio-go/internal/db/seed"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Print("RENIO\n\n")

	env := os.Getenv("RENIO_ENV")
	if env != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("error loading .env file")
		}
		fmt.Printf("! Loaded .env file in %s mode\n", env)
	}

	db.Init()
	if os.Getenv("MIGRATE") == "on" {
		migration.Migrate()
	}
	if os.Getenv("SEED") == "on" {
		seed.RunSeeds()
	}

	Serve()
}
