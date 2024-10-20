package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ad9311/renio-go/internal/app"
	"github.com/ad9311/renio-go/internal/console"
	"github.com/ad9311/renio-go/internal/router"
)

func main() {
	console.AppName()

	if err := app.Init(); err != nil {
		panic(fmt.Sprintf("error initializing app, %s", err.Error()))
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	portF := fmt.Sprintf(":%s", port)
	console.Info(fmt.Sprintf("Listening on http://localhost%s", portF))
	if err := http.ListenAndServe(portF, router.RoutesHandler()); err != nil {
		panic(fmt.Sprintf("there's been an error, %s", err.Error()))
	}
}
