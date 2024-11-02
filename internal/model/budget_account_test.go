package model_test

import (
	"testing"
)

func TestBudgetAccountInsert(t *testing.T) {
	budgetAccount, user := PrepareBudgetAccount(t)

	if budgetAccount.UserID != user.ID {
		t.Errorf("expect budget account with user id %d, got %d", user.ID, budgetAccount.UserID)
	}
}

func TestBudgetAccountSelectByUserID(t *testing.T) {
	budgetAccount, user := PrepareBudgetAccount(t)

	if err := budgetAccount.SelectByUserID(user.ID); err != nil {
		t.Errorf("could not select budget account by user id, %s", err.Error())
	}
}
