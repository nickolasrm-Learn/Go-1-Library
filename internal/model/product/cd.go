package product

import (
	"fmt"

	"github.com/nickolasrm-Learn/Go-2-Library/internal/util"
)

type CD struct {
	id    string
	Title string
	price float64
}

func NewCD(title string, price float64) *CD {
	id := util.NewID()
	return &CD{id: id, Title: title, price: price}
}

func (c *CD) Price() float64 {
	return c.price
}

func (c *CD) ID() string {
	return c.id
}

func (c *CD) String() string {
	return fmt.Sprintf("CD(ID=%v, Title=%q, Price=%v)", c.id, c.Title, c.price)
}
