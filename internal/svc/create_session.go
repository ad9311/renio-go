package svc

import (
	"fmt"
	"time"

	"github.com/ad9311/renio-go/internal/envs"
	"github.com/ad9311/renio-go/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Session struct {
	UserID  int
	JTI     string
	Token   string
	Expires time.Time
}

func CreateSession(userID int) (Session, error) {
	newJWT, err := createJWTToken(userID)
	if err != nil {
		return Session{}, err
	}

	var allowedJWT model.AllowedJWT
	if err = allowedJWT.Insert(newJWT, userID); err != nil {
		return Session{}, nil
	}

	session := Session{
		UserID:  userID,
		JTI:     allowedJWT.JTI,
		Token:   newJWT.Token,
		Expires: allowedJWT.EXP,
	}

	return session, nil
}

// --- Helpers --- //

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
	secret := []byte(envs.GetEnvs().JWTToken)
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
