package makeup

import (
	"errors"
	"time"

	"github.com/rs/zerolog/log"
)

type storage interface {
	GetAllProducts() []Product
	CreateOrder(o Order)
	GetAllOrders() []Order
	GetProductByID(id string) (Product, bool)
}

type Service struct {
	s storage
}

func NewService(s storage) *Service {
	return &Service{s: s}
}

func (s *Service) getPrice(product Product) int {
    currentHour := time.Now().Hour()
    if currentHour >= 18 {
        return int(float64(product.Price) * 1.10) 
    }
    return product.Price
}


func (s *Service) GetProducts() []Product {
	return s.s.GetAllProducts()
}

func (s *Service) PlaceOrder(productID string, quantity int) (Order, error) {
	product, found := s.s.GetProductByID(productID)
	if !found {
		return Order{}, errors.New("product not found")
	}

	totalPrice :=  s.getPrice(product) * quantity
	order := NewOrder(productID, quantity, totalPrice)
	s.s.CreateOrder(order)

	log.Info().Msgf("Order placed: %+v", order)

	return order, nil
}

func (s *Service) GetOrders() []Order {
	return s.s.GetAllOrders()
}
