package user

import (
	"fmt"

	"github.com/nickolasrm-Learn/Go-2-Library/internal/util"
)

type User struct {
	ID     string
	Name   string
	Budget float64
}

// Create a new user with random ID
func New(name string) *User {
	id := util.NewID()
	return &User{ID: id, Name: name}
}

func (u *User) Spend(value float64) error {
	if u.Budget > value {
		u.Budget -= value
		return nil
	}
	return fmt.Errorf("user %q doesn't have enough budget", u.Name)
}

func (u *User) Deposit(value float64) {
	u.Budget += value
}

func (u *User) String() string {
	return fmt.Sprintf("User(name=%q, budget=%v)", u.Name, u.Budget)
}
