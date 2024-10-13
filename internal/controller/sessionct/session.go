package sessionct

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Router(r chi.Router) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", create)
		r.Delete("/", delete)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	var creds credentials
	json.NewDecoder(r.Body).Decode(&creds)
	json.NewEncoder(w).Encode("{}")
}

func delete(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{}")
}
