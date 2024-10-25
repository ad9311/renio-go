package svc

import (
	"fmt"

	"github.com/ad9311/renio-go/internal/eval"
	"github.com/ad9311/renio-go/internal/model"
)

func UserSignUp(signUpData model.SignUpData) (eval.Issues, error) {
	if signUpData.Password != signUpData.PasswordConfirmation {
		return nil, fmt.Errorf("Passwords do not match")
	}

	issues := signUpData.Validate()
	if issues != nil {
		return issues, nil
	}

	var user model.User
	if err := user.Insert(signUpData); err != nil {
		return nil, err
	}

	if err := setUpUserAccounts(user); err != nil {
		return nil, err
	}

	return nil, nil
}

// --- Helpers --- //

func setUpUserAccounts(user model.User) error {
	var budgetAccount model.BudgetAccount
	if err := budgetAccount.Insert(user.ID); err != nil {
		return err
	}

	return nil
}
