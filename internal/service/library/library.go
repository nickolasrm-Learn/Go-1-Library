package library

import (
	"fmt"

	"github.com/nickolasrm-Learn/Go-2-Library/internal/model/product"
	"github.com/nickolasrm-Learn/Go-2-Library/internal/model/purchase"
	"github.com/nickolasrm-Learn/Go-2-Library/internal/model/user"
)

type Library struct {
	Users     map[string]*user.User
	Products  map[string]product.Product
	Purchases map[string]*purchase.Purchase
}

func NewLibrary() *Library {
	return &Library{
		Users:     make(map[string]*user.User),
		Products:  make(map[string]product.Product),
		Purchases: make(map[string]*purchase.Purchase),
	}
}

func (l *Library) AddUser(name string) *user.User {
	usr := user.NewUser(name)
	l.Users[usr.ID] = usr
	return usr
}

func (l *Library) AddProduct(title string, price float64, category string) (product.Product, error) {
	var prod product.Product
	switch category {
	case "CD":
		prod = product.NewCD(title, price)
	case "Book":
		prod = product.NewBook(title, price)
	default:
		return nil, fmt.Errorf("invalid product category %q", category)
	}
	l.Products[prod.ID()] = prod
	return prod, nil
}

func (l *Library) AddPurchase(usr *user.User, prod product.Product) (*purchase.Purchase, error) {
	err := usr.Spend(prod.Price())
	if err != nil {
		return nil, err
	}
	prc := purchase.NewPurchase(usr, prod)
	l.Purchases[prc.ID] = prc
	return prc, nil
}

func (l *Library) ListProducts() []product.Product {
	list := make([]product.Product, len(l.Products))
	i := 0
	for _, v := range l.Products {
		list[i] = v
		i++
	}
	return list
}
