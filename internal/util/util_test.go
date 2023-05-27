package util_test

import (
	"testing"

	"github.com/nickolasrm-Learn/Go-2-Library/internal/util"
)

func TestNewID(t *testing.T) {
	uuid := util.NewID()
	uuid2 := util.NewID()
	if uuid == uuid2 {
		t.Errorf("UUID is returning the same value")
	}
}
