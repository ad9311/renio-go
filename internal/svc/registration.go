package svc

import (
	"fmt"

	"github.com/ad9311/renio-go/internal/model"
)

func SignUpUser(signUpData model.SignUpData) (model.SafeUser, error) {
	var user model.User

	if signUpData.Password != signUpData.PasswordConfirmation {
		return model.SafeUser{}, fmt.Errorf("passwords do not match")
	}

	if err := signUpData.Validate(); err != nil {
		return model.SafeUser{}, err
	}

	if err := user.Insert(signUpData); err != nil {
		return model.SafeUser{}, err
	}

	if err := setUpUserAccounts(user); err != nil {
		return model.SafeUser{}, err
	}

	safeUser := user.GetSafeUser()

	return safeUser, nil
}

// --- Helpers --- //

func setUpUserAccounts(user model.User) error {
	var budgetAccount model.BudgetAccount
	if err := budgetAccount.Insert(user.ID); err != nil {
		return err
	}

	return nil
}
