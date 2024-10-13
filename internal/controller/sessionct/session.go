package sessionct

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ad9311/renio-go/internal/auth"
	"github.com/ad9311/renio-go/internal/controller"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", create)
		r.Delete("/", delete)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	var signInData model.SignInData

	err := json.NewDecoder(r.Body).Decode(&signInData)
	if err != nil {
		controller.WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	var user model.User
	err = user.FindForAuth(signInData.Email)
	if err != nil {
		controller.WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	err = auth.ComparePasswords(user.Password, signInData.Password)
	if err != nil {
		controller.WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	newJWT, err := auth.CreateJWTToken(user.Username)
	if err != nil {
		controller.WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	allowedJWT := model.AllowedJWT{
		JTI:    newJWT.JTI,
		AUD:    newJWT.AUD,
		EXP:    newJWT.EXP,
		UserID: user.ID,
	}
	err = allowedJWT.Insert()
	if err != nil {
		controller.WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", newJWT.Token))
	controller.WriteOK(w, "user signed in successfully", http.StatusCreated)
}

func delete(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{}")
}
