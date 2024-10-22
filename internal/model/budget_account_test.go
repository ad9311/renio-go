package model_test

import (
	"testing"
)

func TestBudgetAccountInsert(t *testing.T) {
	user, err := PrepareUser()
	if err != nil {
		t.Fatalf("could not prepare user for test, %s", err.Error())
	}

	budgetAccount, err := PrepareBudgetAccount(user.ID)
	if err != nil {
		t.Fatalf("could not insert budget account, %s", err.Error())
	}

	if budgetAccount.UserID != user.ID {
		t.Errorf("expect budget account with user id %d, got %d", user.ID, budgetAccount.UserID)
	}
}

func TestBudgetAccountSelectByUserID(t *testing.T) {
	user, err := PrepareUser()
	if err != nil {
		t.Fatalf("could not prepare user for test, %s", err.Error())
	}

	budgetAccount, err := PrepareBudgetAccount(user.ID)
	if err != nil {
		t.Fatalf("could not insert budget account, %s", err.Error())
	}

	if err := budgetAccount.SelectByUserID(user.ID); err != nil {
		t.Errorf("could not select budget account by user id, %s", err.Error())
	}
}
