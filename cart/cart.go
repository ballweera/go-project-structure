package cart

import (
	"github.com/ballweera/go-project-structure/product"
)

type Cart struct {
}

func New() *Cart {
	return &Cart{}
}

func (c *Cart) Add(pd product.Product) error {
	return nil
}
