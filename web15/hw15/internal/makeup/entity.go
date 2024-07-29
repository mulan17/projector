package makeup

import "time"

type Product struct {
	ID          string
	Name        string
	Description string
	Price       int
}

type Order struct {
	ID         string
	ProductID  string
	Quantity   int
	TotalPrice int
	CreatedAt  time.Time
}

func NewProduct(id, name, description string, price int) Product {
	return Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
	}
}

func NewOrder(productID string, quantity int, totalPrice int) Order {
	return Order{
		ID:         time.Now().String(),
		ProductID:  productID,
		Quantity:   quantity,
		TotalPrice: totalPrice,
		CreatedAt:  time.Now(),
	}
}
