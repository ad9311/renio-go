package model_test

import (
	"testing"

	"github.com/ad9311/renio-go/internal/model"
)

func TestInserUser(t *testing.T) {
	signUpData := model.SignUpData{
		Username:             "carlos",
		Name:                 "carlos",
		Password:             "123456789",
		PasswordConfirmation: "123456789",
	}

	var user model.User
	if err := user.Insert(signUpData); err != nil {
		t.Fatalf("failed inserting user: %v", err)
	}

	if user.Username != signUpData.Username {
		t.Errorf("expected user with username % s, got %s", signUpData.Username, user.Username)
	}
}
