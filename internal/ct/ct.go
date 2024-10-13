package ct

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type JWT struct {
	Token string
	JTI   string
	AUD   string
	EXP   time.Time
}

func WriteError(w http.ResponseWriter, errors []string, httpStatus int) error {
	w.WriteHeader(httpStatus)
	return json.NewEncoder(w).Encode(map[string][]string{"errors": errors})
}

func WriteOK(w http.ResponseWriter, data any, httpStatus int) error {
	w.WriteHeader(httpStatus)
	return json.NewEncoder(w).Encode(map[string]any{"data": data})
}

var jwtSecret = []byte(os.Getenv("JWT_KEY"))

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

func comparePasswords(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
