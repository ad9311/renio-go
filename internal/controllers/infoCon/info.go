package infoct

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello, World!",
		Status:  200,
	}

	json.NewEncoder(w).Encode(response)
}
