package model_test

import (
	"testing"
	"time"

	"github.com/ad9311/renio-go/internal/model"
	"github.com/jackc/pgx/v5"
)

func TestAllowJWTInsert(t *testing.T) {
	user := PrepareUser(t)

	jwt := model.JWT{
		JTI: "123456789",
		AUD: "https://renio.dev",
		EXP: time.Now(),
	}

	var aJWT model.AllowedJWT
	if err := aJWT.Insert(jwt, user.ID); err != nil {
		t.Fatalf("could not insert allowed jwt, %s", err.Error())
	}

	if aJWT.JTI != jwt.JTI {
		t.Errorf("expected allowed_jwt with jti to be %s, got %s", jwt.JTI, aJWT.JTI)
	}

	if aJWT.AUD != jwt.AUD {
		t.Errorf("expected allowed_jwt with aud to be %s, got %s", jwt.AUD, aJWT.AUD)
	}

	if !aJWT.EXP.Equal(jwt.EXP) {
		t.Errorf(
			"expected allowed_jwt with exp to be %s, got %s",
			jwt.EXP.String(),
			aJWT.EXP.String(),
		)
	}
}

func TestAllowJWTSelectByJTI(t *testing.T) {
	aJWT := PrepareAllowedJWT(t)

	if err := aJWT.SelectByJTI(aJWT.JTI); err != nil {
		t.Errorf("could not select allowed jwt by jti, %s", err.Error())
	}
}

func TestAllowJWTDelete(t *testing.T) {
	allowedJWT := PrepareAllowedJWT(t)

	if err := allowedJWT.Delete(); err != nil {
		t.Errorf("could not delete allowed jwt, %s", err.Error())
	}

	err := allowedJWT.SelectByJTI(allowedJWT.JTI)
	if err != pgx.ErrNoRows {
		t.Errorf("allowed jwt was not deleted, %s", err.Error())
	}
}
