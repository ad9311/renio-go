package action

import (
	"net/http"
)

// --- Actions --- //

func IndexInfo(w http.ResponseWriter, _ *http.Request) {
	message := "RENIO APP"
	WriteOK(w, message, http.StatusOK)
}
