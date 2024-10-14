package ctrl

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InfoRouter(r chi.Router) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", displayInfo)
	}
}

// Actions //

func displayInfo(w http.ResponseWriter, _ *http.Request) {
	message := "RENIO APP"
	WriteOK(w, message, http.StatusOK)
}
