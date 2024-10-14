package action

import (
	"encoding/json"
	"net/http"
)

func WriteError(w http.ResponseWriter, errors []string, httpStatus int) error {
	w.WriteHeader(httpStatus)
	return json.NewEncoder(w).Encode(map[string][]string{"errors": errors})
}

func WriteOK(w http.ResponseWriter, data any, httpStatus int) error {
	w.WriteHeader(httpStatus)
	return json.NewEncoder(w).Encode(map[string]any{"data": data})
}
