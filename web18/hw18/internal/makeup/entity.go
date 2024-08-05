package makeup

import "time"

type Product struct {
	ID          string `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int `json:"price"`
}

type Order struct {
	ID         string    `json:"ID" db:"id"`
	ProductID  string    `json:"product_id" db:"product_id"`
	Quantity   int       `json:"quantity" db:"quantity"`
	TotalPrice int       `json:"total_price" db:"total_price"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
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
