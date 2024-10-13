package ct

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignInRouter(r chi.Router) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", createSession)
		r.Delete("/", deleteSession)
	}
}

// Actions //

func createSession(w http.ResponseWriter, r *http.Request) {
	var signInData model.SignInData

	err := json.NewDecoder(r.Body).Decode(&signInData)
	if err != nil {
		WriteError(w, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	var user model.User
	err = user.FindForAuth(signInData.Email)
	if err != nil {
		message, status := getFindUserError(err)
		WriteError(w, []string{message}, status)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInData.Password))
	if err != nil {
		message, status := getPasswordError(err)
		WriteError(w, []string{message}, status)
		return
	}

	newJWT, err := createJWTToken(user.Username)
	if err != nil {
		WriteError(w, []string{err.Error()}, http.StatusInternalServerError)
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
		WriteError(w, []string{err.Error()}, http.StatusInternalServerError)
		return
	}

	w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", newJWT.Token))
	WriteOK(w, "user signed in successfully", http.StatusCreated)
}

func deleteSession(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{}")
}

// Helpers //

var jwtSecret = []byte(os.Getenv("JWT_KEY"))

func getFindUserError(err error) (string, int) {
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

func createJWTToken(username string) (JWT, error) {
	var newJWT JWT

	aud := "https://renio.dev"
	iss := "https://api.renio.dev"
	jti := uuid.New().String()
	exp := time.Now().Add(time.Hour * 24 * 7)

	claims := jwt.MapClaims{
		"sub": username,
		"aud": aud,
		"iss": iss,
		"jti": jti,
		"exp": exp.Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return newJWT, err
	}

	newJWT.AUD = aud
	newJWT.EXP = exp
	newJWT.JTI = jti
	newJWT.Token = tokenString

	return newJWT, nil
}
