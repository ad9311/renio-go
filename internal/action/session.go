package action

import (
	"fmt"
	"net/http"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/svc"
	"github.com/ad9311/renio-go/internal/vars"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

// --- Actions --- //

func PostSession(w http.ResponseWriter, r *http.Request) {
	var signInData model.SignInData

	if err := DecodeJSON(r.Body, &signInData); err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	session, err := svc.SignInUser(signInData)
	if err != nil {
		if err == pgx.ErrNoRows || err == bcrypt.ErrMismatchedHashAndPassword {
			err = fmt.Errorf("incorrect email or password")
			WriteError(w, ErrorToSlice(err), http.StatusNotFound)
			return
		}

		WriteError(w, ErrorToSlice(err), http.StatusBadRequest)
		return
	}

	w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", session.Token))
	WriteOK(w, "user signed in successfully", http.StatusCreated)
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	allowedJWT := r.Context().Value(vars.AllowedJWTKey).(model.AllowedJWT)

	if err := svc.SignOutUser(allowedJWT); err != nil {
		WriteError(w, ErrorToSlice(err), http.StatusBadRequest)
		return
	}

	WriteOK(w, "user signed out successfully", http.StatusOK)
}
