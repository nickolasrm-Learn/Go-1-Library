package product_test

import (
	"regexp"
	"testing"

	"github.com/nickolasrm-Learn/Go-2-Library/internal/model/product"
)

func TestCategoryValidateCD(t *testing.T) {
	str := "cd"
	cat := product.Category(str)
	err := cat.Validate()
	if err != nil {
		t.Errorf("CD returned error when validating")
	}
}

func TestCategoryValidateBook(t *testing.T) {
	str := "book"
	cat := product.Category(str)
	err := cat.Validate()
	if err != nil {
		t.Errorf("Book returned error when validating")
	}
}

func TestCategoryValidateInvalid(t *testing.T) {
	str := "invalid"
	cat := product.Category(str)
	err := cat.Validate()
	if err == nil {
		t.Errorf("Invalid category did not return error when validating")
	}
}

func TestNewProduct(t *testing.T) {
	title := "title"
	price := 1.0
	category := product.CD
	prod, err := product.New(title, price, category)
	if err != nil {
		t.Errorf("New returned error when creating product")
	}
	if prod.Title != title {
		t.Errorf("New did not set title correctly")
	}
	if prod.Price != price {
		t.Errorf("New did not set price correctly")
	}
	if prod.Category != category {
		t.Errorf("New did not set category correctly")
	}
}

func TestNewProductInvalidCategory(t *testing.T) {
	title := "title"
	price := 1.0
	category := product.Category("invalid")
	prod, err := product.New(title, price, category)
	if err == nil {
		t.Errorf("New did not return error when creating product with invalid category")
	}
	if prod != nil {
		t.Errorf("New returned product when creating product with invalid category")
	}
}

func TestProductRepresentation(t *testing.T) {
	title := "title"
	price := 1.0
	category := product.CD
	prod, err := product.New(title, price, category)
	if err != nil {
		t.Errorf("New returned error when creating product")
	}
	expected := "Product\\(title=\"title\", price=1, category=\"cd\", id=\".*\"\\)"
	match, _ := regexp.MatchString(expected, prod.String())
	if !match {
		t.Errorf("\"%s\" differs from expected \"%s\"", prod.String(), expected)
	}
}
