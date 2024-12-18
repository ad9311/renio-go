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
		t.Errorf("expected user with username %s, got %s", signUpData.Username, user.Username)
	}

	if user.Username != signUpData.Name {
		t.Errorf("expected user with name %s, got %s", signUpData.Name, user.Name)
	}

	if user.Email != signUpData.Email {
		t.Errorf("expected user with email %s, got %s", signUpData.Email, user.Email)
	}
}

func TestUserSelectByID(t *testing.T) {
	user := PrepareUser(t)

	if err := user.SelectByID(user.ID); err != nil {
		t.Errorf("failed selecting user by id, %s", err.Error())
	}

	err := user.SelectByID(200)
	if err != pgx.ErrNoRows {
		t.Errorf("expected no user got %v from database", user)
	}
}

func TestUserSelectByEmail(t *testing.T) {
	user := PrepareUser(t)

	if err := user.SelectByEmail(user.Email); err != nil {
		t.Errorf("failed selecting user by id, %s", err.Error())
	}

	err := user.SelectByEmail("anon@email.com")
	if err != pgx.ErrNoRows {
		t.Errorf("expected no user got %v from database", user)
	}
}
