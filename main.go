package main

import (
	"fmt"
	"os"

	"github.com/ballweera/go-project-structure/cart"
	"github.com/ballweera/go-project-structure/product"
)

func main() {
	c := cart.New()
	err := c.Add(product.New())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(c)
}
