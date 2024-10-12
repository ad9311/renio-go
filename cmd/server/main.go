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

	if os.Getenv("ENV") == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("error loading .env file")
		}
	}

	db.Init()
	if os.Getenv("MIGRATE") == "auto" {
		db.Migrate()
	}

	Serve()
}
