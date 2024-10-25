package action

import (
	"encoding/json"
	"net/http"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/svc"
)

// --- Actions --- //

func PostUser(w http.ResponseWriter, r *http.Request) {
	var signUpData model.SignUpData
	if err := json.NewDecoder(r.Body).Decode(&signUpData); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	issues, err := svc.UserSignUp(signUpData)
	if issues != nil {
		WriteError(w, issues, http.StatusBadRequest)
		return
	}
	if err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	WriteOK(w, "User created successfully", http.StatusCreated)
}
