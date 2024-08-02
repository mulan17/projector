package makeup

import "time"

type Product struct {
	ID          string `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int `json:"price"`
}

type Order struct {
	ID         string `json:"ID"`
	ProductID  string `json:"product ID"`
	Quantity   int `json:"quantity"`
	TotalPrice int `json:"total price"`
	CreatedAt  time.Time `json:"created time"`
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
