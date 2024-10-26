package svc

import (
	"fmt"

	"github.com/ad9311/renio-go/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type CreatedSession struct{}

func SignInUser(signInData model.SignInData) (Session, error) {
	var user model.User
	if err := user.SelectByEmail(signInData.Email); err != nil {
		return Session{}, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(signInData.Password),
	); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return Session{}, fmt.Errorf("incorrect email or password")
		}
		return Session{}, err
	}

	session, err := CreateSession(user.ID)
	if err != nil {
		return Session{}, err
	}

	return session, nil
}
