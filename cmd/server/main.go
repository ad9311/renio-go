package main

import (
	"fmt"
	"net/http"

	"github.com/ad9311/renio-go/internal/db"
	"github.com/ad9311/renio-go/internal/routes"
)

func main() {
	fmt.Print("RENIO\n\n")

	db.Init()

	err := http.ListenAndServe(":8080", routes.Router())
	if err != nil {
		fmt.Printf("there's been an error: %s", err.Error())
	}
}
