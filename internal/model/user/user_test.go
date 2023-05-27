package user_test

import (
	"testing"

	"github.com/nickolasrm-Learn/Go-2-Library/internal/model/user"
)

func TestNewUser(t *testing.T) {
	name := "name"
	user := user.New(name)
	if user.Name != name {
		t.Errorf("New did not set name correctly")
	}
	if user.ID == "" {
		t.Errorf("New did not set ID correctly")
	}
}

func TestUserDeposit(t *testing.T) {
	user := user.New("name")
	user.Deposit(1.0)
	if user.Budget != 1.0 {
		t.Errorf("Deposit did not set budget correctly")
	}
	user.Deposit(1.0)
	if user.Budget != 2.0 {
		t.Errorf("Deposit did not set budget correctly")
	}
}

func TestUserSpend(t *testing.T) {
	usr := user.New("name")
	usr.Deposit(1.0)
	err := usr.Spend(1.0)
	if err != nil {
		t.Errorf("Spend returned error when spending")
	}
	if usr.Budget != 0.0 {
		t.Errorf("Spend did not set budget correctly")
	}
	err = usr.Spend(1.0)
	if err == nil {
		t.Errorf("Spend did not return error when spending more than budget")
	}
	if usr.Budget != 0.0 {
		t.Errorf("Spend set budget incorrectly")
	}
}

func TestUserStringer(t *testing.T) {
	usr := user.New("name")
	usr.Deposit(1.0)
	if usr.String() != "User(name=\"name\", budget=1)" {
		t.Errorf("Stringer returned unexpected string")
	}
}
