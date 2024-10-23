package action

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ad9311/renio-go/internal/console"
)

func WriteError(w http.ResponseWriter, errors []string, httpStatus int) {
	w.WriteHeader(httpStatus)
	if err := json.NewEncoder(w).Encode(map[string][]string{"errors": errors}); err != nil {
		console.Error(fmt.Sprintf("could not write response, %s", err.Error()))
	}
}

func WriteOK(w http.ResponseWriter, data any, httpStatus int) {
	w.WriteHeader(httpStatus)
	if err := json.NewEncoder(w).Encode(map[string]any{"data": data}); err != nil {
		console.Error(fmt.Sprintf("could not write response, %s", err.Error()))
	}
}
