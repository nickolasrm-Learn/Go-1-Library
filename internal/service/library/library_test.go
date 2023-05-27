package library_test

import (
	"testing"

	"github.com/nickolasrm-Learn/Go-2-Library/internal/model/product"
	"github.com/nickolasrm-Learn/Go-2-Library/internal/model/purchase"
	"github.com/nickolasrm-Learn/Go-2-Library/internal/model/user"
	"github.com/nickolasrm-Learn/Go-2-Library/internal/service/library"
)

func EmptyLibrary() *library.Library {
	return &library.Library{
		Users:     make(map[string]*user.User),
		Products:  make(map[string]*product.Product),
		Purchases: make(map[string]*purchase.Purchase),
	}
}

func TestLibraryAddProduct(t *testing.T) {
	library := EmptyLibrary()
	prod, _ := library.AddProduct("test", 0.5, "book")
	prod_ref, ok := library.Products[prod.ID]
	if !ok || prod_ref != prod {
		t.Errorf("Product was not added to library")
	}
}

func TestLibraryAddInvalidProduct(t *testing.T) {
	library := EmptyLibrary()
	_, err := library.AddProduct("test", 0.5, "abc")
	if err == nil {
		t.Errorf("Invalid product was added to library")
	}
}

func TestLibraryAddUser(t *testing.T) {
	library := EmptyLibrary()
	user := library.AddUser("test")
	user_ref, ok := library.Users[user.ID]
	if !ok || user_ref != user {
		t.Errorf("User was not added to library")
	}
}

func TestLibraryGetInexistentUser(t *testing.T) {
	library := EmptyLibrary()
	_, err := library.GetUser("abc")
	if err == nil {
		t.Errorf("Inexistent user was retrieved from library")
	}
}

func TestLibraryGetProduct(t *testing.T) {
	library := EmptyLibrary()
	prod, _ := library.AddProduct("test", 0.5, "book")
	prod_ref, err := library.GetProduct(prod.ID)
	if err != nil || prod_ref != prod {
		t.Errorf("Product was not retrieved from library")
	}
}

func TestLibraryGetInexistentProduct(t *testing.T) {
	library := EmptyLibrary()
	_, err := library.GetProduct("abc")
	if err == nil {
		t.Errorf("Inexistent product was retrieved from library")
	}
}

func TestLibraryListProducts(t *testing.T) {
	library := EmptyLibrary()
	prod, _ := library.AddProduct("test", 0.5, "book")
	prod2, _ := library.AddProduct("test2", 0.5, "book")
	prods := library.ListProducts()
	if len(prods) != 2 || prods[0] != prod || prods[1] != prod2 {
		t.Errorf("Products were not listed correctly")
	}
}

func TestLibraryDeposit(t *testing.T) {
	library := EmptyLibrary()
	user := library.AddUser("test")
	user.Deposit(10)
	if user.Budget != 10 {
		t.Errorf("Deposit was not made correctly")
	}
}

func TestLibraryBuy(t *testing.T) {
	library := EmptyLibrary()
	user := library.AddUser("test")
	prod, _ := library.AddProduct("test", 0.5, "book")
	library.Deposit(user, 1.0)
	purchase, _ := library.Buy(user, prod)
	if purchase.User != user.ID || purchase.Product != prod.ID {
		t.Errorf("Purchase was not made correctly")
	}
	if user.Budget != 0.5 {
		t.Errorf("Purchase was not made correctly")
	}
}

func TestLibraryBuyNoBudget(t *testing.T) {
	library := EmptyLibrary()
	user := library.AddUser("test")
	prod, _ := library.AddProduct("test", 0.5, "book")
	_, err := library.Buy(user, prod)
	if err == nil {
		t.Errorf("Purchase was made without budget")
	}
}

func TestLibraryListPurchases(t *testing.T) {
	library := EmptyLibrary()
	user := library.AddUser("test")
	prod, _ := library.AddProduct("test", 0.5, "book")
	library.Deposit(user, 1.0)
	library.Buy(user, prod)
	purchases := library.ListPurchases(user)
	if len(purchases) != 1 || purchases[0].User != user.ID || purchases[0].Product != prod.ID {
		t.Errorf("Purchases were not listed correctly")
	}
}
