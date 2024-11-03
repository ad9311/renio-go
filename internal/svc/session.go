package svc

import (
	"github.com/ad9311/renio-go/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func SignInUser(signInData model.SignInData) (model.SafeUser, error) {
	var user model.User

	if err := user.SelectByEmail(signInData.Email); err != nil {
		return model.SafeUser{}, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(signInData.Password),
	); err != nil {
		return model.SafeUser{}, err
	}

	safeUser := user.GetSafeUser()

	return safeUser, nil
}

func SignOutUser() error {
	return nil
}

func CreateSession(userID int) error {
	return nil
}
