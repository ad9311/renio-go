package svc

import (
	"github.com/ad9311/renio-go/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func SignInUser(signInData model.SignInData) (model.User, error) {
	var user model.User

	if err := user.SelectByEmail(signInData.Email); err != nil {
		return user, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(signInData.Password),
	); err != nil {
		return user, err
	}

	return user, nil
}
