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
	var user model.User
	if err := user.SelectByID(1); err != nil {
		t.Fatalf("failed selecting user by id, %s", err.Error())
	}

	err := user.SelectByID(2)
	if err != pgx.ErrNoRows {
		t.Errorf("expected no user got %v from database", user)
	}
}

func TestUserSelectByEmail(t *testing.T) {
	var user model.User
	if err := user.SelectByEmail("carlos@email.com"); err != nil {
		t.Fatalf("failed selecting user by id, %s", err.Error())
	}

	err := user.SelectByEmail("andres@email.com")
	if err != pgx.ErrNoRows {
		t.Errorf("expected no user got %v from database", user)
	}
}
