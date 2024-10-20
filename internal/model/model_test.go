package model_test

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/ad9311/renio-go/internal/app"
	"github.com/ad9311/renio-go/internal/db/migration"
	"github.com/ad9311/renio-go/internal/model"
)

func TestMain(m *testing.M) {
	if err := app.Init(); err != nil {
		log.Fatal(err.Error())
	}

	code := m.Run()
	cleanUp()
	os.Exit(code)
}

func cleanUp() {
	if err := migration.Reset(); err != nil {
		panic(err.Error())
	}
}

func prepareUser() (model.User, error) {
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
