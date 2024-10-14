package action

import (
	"encoding/json"
	"net/http"

	"github.com/ad9311/renio-go/internal/model"
)

// Actions //

func PostUser(w http.ResponseWriter, r *http.Request) {
	var signUpData model.SignUpData
	if err := json.NewDecoder(r.Body).Decode(&signUpData); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	if signUpData.Password != signUpData.PasswordConfirmation {
		WriteError(w, []string{"passwords don't match"}, http.StatusBadRequest)
		return
	}

	var user model.User
	if err := user.Create(signUpData); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	if err := user.SetUpAccounts(); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	WriteOK(w, "user created successfully", http.StatusCreated)
}
