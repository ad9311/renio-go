package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ad9311/renio-go/internal/app"
	"github.com/ad9311/renio-go/internal/conf"
	"github.com/ad9311/renio-go/internal/router"
)

func main() {
	if err := app.Init(); err != nil {
		log.Fatalf("error initializing app, %v", err)
	}

	port := fmt.Sprintf(":%s", conf.GetEnv().Port)
	session := conf.GetSession()
	routesHandler := router.RoutesHandler()

	err := http.ListenAndServe(port, session.LoadAndSave(routesHandler))
	if err != nil {
		log.Fatalf("there's been an error, %v", err)
	}
}
