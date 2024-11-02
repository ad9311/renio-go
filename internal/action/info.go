package action

import (
	"net/http"
)

// --- Actions --- //

func IndexInfo(w http.ResponseWriter, _ *http.Request) {
	dataResponse := DataResponse{
		Content: Content{
			"appName":   "RENIO APP",
			"createdBy": "Ángel Díaz",
		},
	}
	WriteOK(w, dataResponse)
}
