package purchase_test

import (
	"testing"

	"github.com/nickolasrm-Learn/Go-2-Library/internal/model/product"
	"github.com/nickolasrm-Learn/Go-2-Library/internal/model/purchase"
	"github.com/nickolasrm-Learn/Go-2-Library/internal/model/user"
)

func TestNewPurchase(t *testing.T) {
	// Create a new user and product
	user := user.New("John Doe")
	user.Deposit(1000)
	prod, _ := product.New("The Art of Computer Programming", 1000, "book")
	purchase := purchase.New(user, prod)
	if purchase.User != user.ID {
		t.Errorf("New did not set user correctly")
	}
	if purchase.Product != prod.ID {
		t.Errorf("New did not set product correctly")
	}
}
