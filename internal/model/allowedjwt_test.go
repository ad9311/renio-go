package model_test

import (
	"testing"
	"time"

	"github.com/ad9311/renio-go/internal/model"
)

func TestAllowJWTInsert(t *testing.T) {
	user, err := prepareUser()
	if err != nil {
		t.Fatalf("could not prepare user for test, %s", err.Error())
	}

	jwt := model.JWT{
		JTI: "123456789",
		AUD: "https://renio.dev",
		EXP: time.Now(),
	}

	var aJWT model.AllowedJWT
	if err := aJWT.Insert(jwt, user.ID); err != nil {
		t.Fatalf("could not insert allowed jwt, %s", err.Error())
	}
}

func TestAllowJWTSelectByJTI(t *testing.T) {
	var aJWT model.AllowedJWT
	if err := aJWT.SelectByJTI("123456789"); err != nil {
		t.Fatalf("could not select allowed jwt by jti, %s", err.Error())
	}
}
