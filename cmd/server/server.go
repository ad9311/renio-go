package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ad9311/renio-go/internal/console"
	"github.com/ad9311/renio-go/internal/router"
)

func Serve() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	console.Info(fmt.Sprintf("Listening on http://localhost:%s", port))

	portF := fmt.Sprintf(":%s", port)
	if err := http.ListenAndServe(portF, router.RoutesHandler()); err != nil {
		console.Fatal(fmt.Sprintf("there's been an error, %s", err.Error()))
	}
}
