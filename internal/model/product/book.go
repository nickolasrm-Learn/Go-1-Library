package product

import (
	"fmt"

	"github.com/nickolasrm-Learn/Go-2-Library/internal/util"
)

type Book struct {
	id    string
	Title string
	price float64
}

func NewBook(title string, price float64) *Book {
	id := util.NewID()
	return &Book{id: id, Title: title, price: price}
}

func (b *Book) Price() float64 {
	return b.price
}

func (b *Book) ID() string {
	return b.id
}

func (b *Book) String() string {
	return fmt.Sprintf("Book(ID=%v, Title=%q, Price=%v)", b.id, b.Title, b.price)
}
