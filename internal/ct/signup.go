package ct

import (
	"encoding/json"
	"net/http"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/go-chi/chi/v5"
)

func SignUpRouter(r chi.Router) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", createUser)
	}
}

// Actions //

func createUser(w http.ResponseWriter, r *http.Request) {
	var signUpData model.SignUpData
	err := json.NewDecoder(r.Body).Decode(&signUpData)
	if err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	if signUpData.Password != signUpData.PasswordConfirmation {
		WriteError(w, []string{"passwords don't match"}, http.StatusBadRequest)
		return
	}

	var user model.User
	err = user.Create(signUpData)
	if err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	WriteOK(w, "user created successfully", http.StatusCreated)
}
