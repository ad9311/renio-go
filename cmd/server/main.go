package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ad9311/renio-go/internal/db"
	"github.com/ad9311/renio-go/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Print("RENIO\n\n")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	db.Init()

	err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), routes.Router())
	if err != nil {
		fmt.Printf("there's been an error: %s", err.Error())
	}
}
