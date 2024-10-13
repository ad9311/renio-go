package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_KEY"))

type JWT struct {
	Token string
	JTI   string
	AUD   string
	EXP   time.Time
}

func CreateJWTToken(username string) (JWT, error) {
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

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePasswords(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
