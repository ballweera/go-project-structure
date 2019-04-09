package cart

type Cart struct {
}

func New() *Cart {
	return &Cart{}
}

type Product struct {
}

func (c *Cart) Add(p Product) error {
	return nil
}
