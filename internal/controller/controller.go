package controller

import (
	"encoding/json"
	"net/http"
)

func WriteError(w http.ResponseWriter, errors []string, httpStatus int) {
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(map[string][]string{"errors": errors})
}

func WriteOK(w http.ResponseWriter, data any, httpStatus int) {
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(map[string]any{"data": data})
}
