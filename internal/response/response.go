package response

import (
	"encoding/json"
	"net/http"
)

func InvalidJSON(w http.ResponseWriter, error string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error": error,
	})
}
