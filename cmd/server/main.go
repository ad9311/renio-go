package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ad9311/renio-go/internal/db"
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
	if os.Getenv("MIGRATE") == "auto" {
		db.Migrate()
	}

	Serve()
}
