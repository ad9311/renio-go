package routes

import (
	"net/http"

	infoct "github.com/ad9311/renio-go/internal/controllers/infoct"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() http.Handler {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(headerRouter)

	r.Route("/", func(r chi.Router) {
		// Info
		r.Get("/info", infoct.Index)
	})

	return r
}

func headerRouter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
