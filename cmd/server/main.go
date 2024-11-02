package main

import (
	"fmt"
	"net/http"

	"github.com/ad9311/renio-go/internal/app"
	"github.com/ad9311/renio-go/internal/cnsl"
	"github.com/ad9311/renio-go/internal/conf"
	"github.com/ad9311/renio-go/internal/router"
)

func main() {
	cnsl.AppName()

	if err := app.Init(); err != nil {
		panic(fmt.Sprintf("error initializing app, %s", err.Error()))
	}

	portF := fmt.Sprintf(":%s", conf.GetEnv().Port)
	cnsl.Info(fmt.Sprintf("Listening on http://localhost%s", portF))
	if err := http.ListenAndServe(portF, router.RoutesHandler()); err != nil {
		panic(fmt.Sprintf("there's been an error, %s", err.Error()))
	}
}
