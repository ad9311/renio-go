package svc

import (
	"fmt"

	"github.com/ad9311/renio-go/internal/model"
)

func SignUpUser(signUpData model.SignUpData) error {
	if signUpData.Password != signUpData.PasswordConfirmation {
		return fmt.Errorf("Passwords do not match")
	}

	if err := signUpData.Validate(); err != nil {
		return err
	}

	var user model.User
	if err := user.Insert(signUpData); err != nil {
		return err
	}

	if err := setUpUserAccounts(user); err != nil {
		return err
	}

	return nil
}

// --- Helpers --- //

func setUpUserAccounts(user model.User) error {
	var budgetAccount model.BudgetAccount
	if err := budgetAccount.Insert(user.ID); err != nil {
		return err
	}

	return nil
}
