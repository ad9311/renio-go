package infoct

import (
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{}")
}
