package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ad9311/renio-go/internal/router"
)

func Serve() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("! Listening on http://localhost:%s\n\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router.RoutesHandler())
	if err != nil {
		log.Fatalf("there's been an error: %s", err.Error())
	}
}
