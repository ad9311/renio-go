package route

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/controller/infoct"
	"github.com/ad9311/renio-go/internal/controller/sessionct"
	reniomiddleware "github.com/ad9311/renio-go/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() http.Handler {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(reniomiddleware.HeaderRouter)
	r.Use(reniomiddleware.JSONValidator)

	r.Route("/", func(r chi.Router) {
		// Info
		r.Get("/info", infoct.Index)

		// Sessions
		r.Route("/", sessionct.Router(r))
	})

	return r
}
