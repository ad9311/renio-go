package model_test

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"

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

func PrepareUser(t *testing.T) model.User {
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
		t.Fatalf("could not prepare user for test, %s", err.Error())
	}

	return user
}

func PrepareBudgetAccount(t *testing.T) (model.BudgetAccount, model.User) {
	user := PrepareUser(t)

	var budgetAccount model.BudgetAccount
	if err := budgetAccount.Insert(user.ID); err != nil {
		t.Fatalf("could not prepare budget account for test, %s", err.Error())
	}

	return budgetAccount, user
}

func PrepareBudget(t *testing.T) model.Budget {
	budgetAccount, _ := PrepareBudgetAccount(t)

	var budget model.Budget
	if err := budget.Insert(budgetAccount.ID); err != nil {
		t.Fatalf("could not prepare budget for test, %s", err.Error())
	}

	return budget
}
