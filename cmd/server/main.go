package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ad9311/renio-go/internal/app"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/router"
)

func main() {
	fmt.Print("RENIO\n\n")

	if err := app.InitEnv(); err != nil {
		log.Fatalf("there's been an error initializing the environment, %v", err)
	}

	if err := app.InitDB(); err != nil {
		log.Fatalf("there's been an error initializing the database, %v", err)
	}

	if app.GetEnv().AppEnv != app.Test {
		app.InitSessionManager()

		if _, err := app.BuildTemplateCache(); err != nil {
			log.Fatalf("there's been an error initializing the template cache, %v", err)
		}
	}

	if app.GetEnv().AppEnv == app.Production {
		if err := app.Migrate(); err != nil {
			log.Fatalf("there's been an error migrating the database, %v", err)
		}
	}

	if app.GetEnv().Seed {
		if err := model.RunSeeds(); err != nil {
			log.Fatalf("there's been an error seeding the database, %v", err)
		}
	}

	model.RegisterModels()

	port := fmt.Sprintf(":%s", app.GetEnv().Port)
	if err := http.ListenAndServe(port, router.RoutesHandler()); err != nil {
		log.Fatalf("there's been an error booting the server, %v", err)
	}
}
