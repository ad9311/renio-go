package lib

import (
	"os"
	"time"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret = []byte(os.Getenv("JWT_KEY"))

func CreateJWTToken(username string) (model.NewJWT, error) {
	var newJWT model.NewJWT

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
