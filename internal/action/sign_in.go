package action

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

// --- Actions --- //

func PostSession(w http.ResponseWriter, r *http.Request) {
	var signInData model.SignInData

	err := json.NewDecoder(r.Body).Decode(&signInData)
	if err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	var user model.User
	err = user.SelectForAuth(signInData.Email)
	if err != nil {
		message, status := getSelectUserError(err)
		WriteError(w, []string{message}, status)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInData.Password))
	if err != nil {
		message, status := getPasswordError(err)
		WriteError(w, []string{message}, status)
		return
	}

	newJWT, err := createJWTToken(user.ID)
	if err != nil {
		WriteError(w, []string{err.Error()}, http.StatusInternalServerError)
		return
	}

	var allowedJWT model.AllowedJWT
	err = allowedJWT.Insert(newJWT, user.ID)
	if err != nil {
		WriteError(w, []string{err.Error()}, http.StatusInternalServerError)
		return
	}

	w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", newJWT.Token))
	WriteOK(w, "user signed in successfully", http.StatusCreated)
}

// --- Helpers --- //

func getSelectUserError(err error) (string, int) {
	var message string
	var status int

	if err == pgx.ErrNoRows {
		message = "wrong password or email"
		status = http.StatusUnauthorized
	} else {
		message = err.Error()
		status = http.StatusBadRequest
	}

	return message, status
}

func getPasswordError(err error) (string, int) {
	var message string
	var status int

	if err == bcrypt.ErrMismatchedHashAndPassword {
		message = "wrong password or email"
		status = http.StatusUnauthorized
	} else {
		message = err.Error()
		status = http.StatusBadRequest
	}

	return message, status
}

func createJWTToken(userID int) (model.JWT, error) {
	jti := uuid.New().String()
	exp := time.Now().Add(time.Hour * 24 * 7)

	claims := jwt.MapClaims{
		"sub": fmt.Sprintf("%d", userID),
		"jti": jti,
		"exp": exp.Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv("JWT_KEY"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return model.JWT{}, err
	}

	newJWT := model.JWT{
		JTI:   jti,
		EXP:   exp,
		Token: tokenString,
	}

	return newJWT, nil
}
