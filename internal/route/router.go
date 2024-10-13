package route

import (
	"net/http"

	"github.com/ad9311/renio-go/internal/controller/infoct"
	"github.com/ad9311/renio-go/internal/controller/registrationct"
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

	r.Route("/", func(r chi.Router) {
		// Info
		r.Get("/info", infoct.Index)

		// Auth
		r.Route("/auth", func(r chi.Router) {
			// Sessions
			r.Route("/sign-in", sessionct.Router(r))

			// Registrations
			r.Route("/sign-up", registrationct.Router(r))
		})
	})

	return r
}
