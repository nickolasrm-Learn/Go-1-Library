package product

import (
	"fmt"

	"github.com/nickolasrm-Learn/Go-2-Library/internal/util"
)

type Category string

const (
	CD   Category = "cd"
	Book Category = "book"
)

func (c Category) Validate() error {
	switch c {
	case CD, Book:
		break
	default:
		return fmt.Errorf("invalid category %q", c)
	}
	return nil
}

type Product struct {
	ID       string
	Title    string
	Price    float64
	Category Category
}

func New(title string, price float64, category Category) (*Product, error) {
	id := util.NewID()
	err := category.Validate()
	if err != nil {
		return nil, err
	}
	return &Product{id, title, price, category}, nil
}

func (p *Product) String() string {
	return fmt.Sprintf("Product(title=%q, price=%v, category=%q, id=%q)", p.Title, p.Price, p.Category, p.ID)
}
