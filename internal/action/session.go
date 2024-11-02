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

	errResponse := ErrorResponse{}

	if err := DecodeJSON(r.Body, &signInData); err != nil {
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}

	session, err := svc.SignInUser(signInData)
	if err == pgx.ErrNoRows || err == bcrypt.ErrMismatchedHashAndPassword {
		err = fmt.Errorf("incorrect email or password")
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}
	if err != nil {
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}

	dataResponse := DataResponse{
		Content: Content{
			"message": "user signed in successfully",
		},
	}

	w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", session.Token))
	WriteOK(w, dataResponse)
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	allowedJWT := r.Context().Value(vars.AllowedJWTKey).(model.AllowedJWT)

	errResponse := ErrorResponse{}
	if err := svc.SignOutUser(allowedJWT); err != nil {
		errResponse.Append(err)
		WriteError(w, errResponse)
		return
	}

	dataResponse := DataResponse{
		Content: Content{
			"message": "user signed out successfully",
		},
	}

	WriteOK(w, dataResponse)
}
