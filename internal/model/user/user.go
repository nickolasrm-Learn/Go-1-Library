package user

import (
	"fmt"

	"github.com/nickolasrm-Learn/Go-2-Library/internal/util"
)

// User represents a user of the library
type User struct {
	ID     string
	Name   string
	Budget float64
}

// New creates a new user with random ID
func New(name string) *User {
	id := util.NewID()
	return &User{ID: id, Name: name}
}

// Spend subtracts value from user budget
//
// Returns an error if user doesn't have enough budget
func (u *User) Spend(value float64) error {
	if u.Budget >= value {
		u.Budget -= value
		return nil
	}
	return fmt.Errorf("user %q doesn't have enough budget", u.Name)
}

// Deposit adds value to user budget
func (u *User) Deposit(value float64) {
	u.Budget += value
}

// String returns a string representation of the user
func (u *User) String() string {
	return fmt.Sprintf("User(name=%q, budget=%v)", u.Name, u.Budget)
}
