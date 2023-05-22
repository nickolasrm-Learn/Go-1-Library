package util

import (
	"github.com/google/uuid"
)

// NewID returns a new random ID
func NewID() string {
	return uuid.NewString()
}
