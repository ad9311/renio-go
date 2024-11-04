package svc

import (
	"fmt"

	"github.com/ad9311/renio-go/internal/model"
)

func SignUpUser(signUpData model.SignUpData) (model.User, error) {
	var user model.User

	if signUpData.Password != signUpData.PasswordConfirmation {
		return user, fmt.Errorf("passwords do not match")
	}

	if err := signUpData.Validate(); err != nil {
		return user, err
	}

	if err := user.Insert(signUpData); err != nil {
		return user, err
	}

	if err := setUpUserAccounts(user); err != nil {
		return user, err
	}

	return user, nil
}

// --- Helpers --- //

func setUpUserAccounts(user model.User) error {
	var budgetAccount model.BudgetAccount
	if err := budgetAccount.Insert(user.ID); err != nil {
		return err
	}

	return nil
}
