package product

import (
	"fmt"

	"github.com/nickolasrm-Learn/Go-2-Library/internal/util"
)

// Category represents the string category of a product
type Category string

const (
	// CD represents the CD category
	CD Category = "cd"
	// Book represents the Book category
	Book Category = "book"
)

// Validate checks if the category is valid
func (c Category) Validate() error {
	switch c {
	case CD, Book:
		break
	default:
		return fmt.Errorf("invalid category %q", c)
	}
	return nil
}

// Product represents a product in the library
type Product struct {
	ID       string
	Title    string
	Price    float64
	Category Category
}

// New creates a new product
//
// It returns an error if the category is invalid
func New(title string, price float64, category Category) (*Product, error) {
	id := util.NewID()
	err := category.Validate()
	if err != nil {
		return nil, err
	}
	return &Product{id, title, price, category}, nil
}

// String returns a string representation of the product
func (p *Product) String() string {
	return fmt.Sprintf("Product(title=%q, price=%v, category=%q, id=%q)", p.Title, p.Price, p.Category, p.ID)
}
