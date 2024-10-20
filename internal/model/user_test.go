package model_test

import (
	"testing"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/jackc/pgx/v5"
)

func TestUserInsert(t *testing.T) {
	signUpData := model.SignUpData{
		Username:             "carlos",
		Name:                 "carlos",
		Email:                "carlos@email.com",
		Password:             "123456789",
		PasswordConfirmation: "123456789",
	}

	var user model.User
	if err := user.Insert(signUpData); err != nil {
		t.Fatalf("failed inserting user, %s", err.Error())
	}

	if user.Username != signUpData.Username {
		t.Errorf("expected user with username % s, got %s", signUpData.Username, user.Username)
	}

	if user.Username != signUpData.Name {
		t.Errorf("expected user with name % s, got %s", signUpData.Name, user.Name)
	}
}

func TestUserSelectByID(t *testing.T) {
	user, err := prepareUser()
	if err != nil {
		t.Fatalf("could not prepare user for test, %s", err.Error())
	}

	if err := user.SelectByID(user.ID); err != nil {
		t.Fatalf("failed selecting user by id, %s", err.Error())
	}

	err = user.SelectByID(20)
	if err != pgx.ErrNoRows {
		t.Errorf("expected no user got %v from database", user)
	}
}

func TestUserSelectByEmail(t *testing.T) {
	user, err := prepareUser()
	if err != nil {
		t.Fatalf("could not prepare user for test, %s", err.Error())
	}

	if err := user.SelectByEmail(user.Email); err != nil {
		t.Fatalf("failed selecting user by id, %s", err.Error())
	}

	err = user.SelectByEmail("anon@email.com")
	if err != pgx.ErrNoRows {
		t.Errorf("expected no user got %v from database", user)
	}
}
