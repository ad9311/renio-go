package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ad9311/renio-go/internal/route"
)

func Serve() {
	port := os.Getenv("PORT")

	fmt.Printf("! Listening on http://localhost:%s\n\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), route.Router())
	if err != nil {
		log.Fatalf("there's been an error: %s", err.Error())
	}
}
