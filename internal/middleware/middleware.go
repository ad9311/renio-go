package middleware

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ad9311/renio-go/internal/response"
)

func JSONValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only validate for requests with JSON body
		if r.Header.Get("Content-Type") == "application/json" && r.ContentLength > 0 {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				response.InvalidJSON(w, "Failed to read request body", http.StatusBadRequest)
				return
			}

			var js json.RawMessage
			if err := json.Unmarshal(body, &js); err != nil {
				response.InvalidJSON(w, "Malformed JSON", http.StatusBadRequest)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func HeaderRouter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
