package action

import (
	"fmt"
	"net/http"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/ad9311/renio-go/internal/svc"
	"github.com/jackc/pgx/v5"
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
		if err == pgx.ErrNoRows {
			err = fmt.Errorf("incorrect email or password")
			WriteError(w, ErrorToSlice(err), http.StatusNotFound)
			return
		}

		WriteError(w, ErrorToSlice(err), http.StatusNotFound)
		return
	}

	w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", session.Token))
	WriteOK(w, "user signed in successfully", http.StatusCreated)
}
