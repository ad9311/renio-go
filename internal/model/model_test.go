package model_test

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/ad9311/renio-go/internal/app"
	"github.com/ad9311/renio-go/internal/model"
)

func TestMain(m *testing.M) {
	if err := app.Init(); err != nil {
		log.Fatal(err.Error())
	}

	code := m.Run()
	os.Exit(code)
}

func PrepareUser() (model.User, error) {
	str := fmt.Sprintf("%d", rand.Int())
	signUpData := model.SignUpData{
		Username:             str,
		Name:                 "Jon Doe",
		Email:                fmt.Sprintf("%s@mail.com", str),
		Password:             "123456789",
		PasswordConfirmation: "123456789",
	}

	var user model.User
	if err := user.Insert(signUpData); err != nil {
		return user, err
	}

	return user, nil
}

func PrepareAllowedJWT(userID int) (model.AllowedJWT, error) {
	str := fmt.Sprintf("%d", rand.Int())
	jwt := model.JWT{
		Token: str,
		JTI:   str,
		AUD:   "http://localhost:3000",
		EXP:   time.Now(),
	}
	var aJWT model.AllowedJWT
	if err := aJWT.Insert(jwt, userID); err != nil {
		return aJWT, err
	}

	return aJWT, nil
}

func PrepareBudgetAccount(userID int) (model.BudgetAccount, error) {
	var budgetAccount model.BudgetAccount
	if err := budgetAccount.Insert(userID); err != nil {
		return budgetAccount, err
	}

	return budgetAccount, nil
}
