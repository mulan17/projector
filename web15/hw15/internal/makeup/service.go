package makeup

import (
	"errors"

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

func (s *Service) GetProducts() []Product {
	return s.s.GetAllProducts()
}

func (s *Service) PlaceOrder(productID string, quantity int) (Order, error) {
	product, found := s.s.GetProductByID(productID)
	if !found {
		return Order{}, errors.New("product not found")
	}

	totalPrice := product.Price * quantity
	order := NewOrder(productID, quantity, totalPrice)
	s.s.CreateOrder(order)

	log.Info().Msgf("Order placed: %+v", order)

	return order, nil
}

func (s *Service) GetOrders() []Order {
	return s.s.GetAllOrders()
}
