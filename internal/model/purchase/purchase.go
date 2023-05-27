package purchase

import (
	"github.com/nickolasrm-Learn/Go-2-Library/internal/model/product"
	"github.com/nickolasrm-Learn/Go-2-Library/internal/model/user"
	"github.com/nickolasrm-Learn/Go-2-Library/internal/util"
)

type Purchase struct {
	ID      string
	User    string
	Product string
}

func New(user *user.User, prod *product.Product) *Purchase {
	return &Purchase{ID: util.NewID(), User: user.ID, Product: prod.ID}
}
