package util

import (
	"github.com/google/uuid"
)

// NewID returns a new random ID string
func NewID() string {
	return uuid.NewString()
}
