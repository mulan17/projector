package makeup

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

type storage interface {
    GetAllProducts() ([]Product, error)
    CreateProduct(p Product) error
    CreateOrder(o Order) error
    GetAllOrders() ([]Order, error)
    GetProductByID(id string) (Product, error)
}

type Service struct {
    s storage
}

func NewService(s storage) *Service {
    return &Service{s: s}
}

func (s *Service) CreateProduct(p Product) error {
    if err := s.s.CreateProduct(p); err != nil {
        return fmt.Errorf("creating product: %w", err)
    }
    return nil
}

func (s *Service) GetAllProducts() ([]Product, error) {
    products, err := s.s.GetAllProducts()
    if err != nil {
        return nil, fmt.Errorf("getting all products: %w", err)
    }
    return products, nil
}

func (s *Service) GetProductByID(id string) (Product, error) {
    product, err := s.s.GetProductByID(id)
    if err != nil {
        return Product{}, fmt.Errorf("getting product by ID: %w", err)
    }
    return product, nil
}

func (s *Service) CreateOrder(o Order) error {
    if err := s.s.CreateOrder(o); err != nil {
        return fmt.Errorf("creating order: %w", err)
    }
    return nil
}

func (s *Service) GetAllOrders() ([]Order, error) {
    orders, err := s.s.GetAllOrders()
    if err != nil {
        return nil, fmt.Errorf("getting all orders: %w", err)
    }
    return orders, nil
}

func (s *Service) PlaceOrder(productID string, quantity int) (Order, error) {
    product, err := s.s.GetProductByID(productID)
    if err != nil {
        return Order{}, fmt.Errorf("getting product by ID: %w", err)
    }

    totalPrice := s.getPrice(product) * quantity
    order := NewOrder(productID, quantity, totalPrice)
    if err := s.s.CreateOrder(order); err != nil {
        return Order{}, fmt.Errorf("creating order: %w", err)
    }

    log.Info().Msgf("Order placed: %+v", order)

    return order, nil
}

func (s *Service) getPrice(product Product) int {
    currentHour := time.Now().Hour()
    if currentHour >= 18 {
        return int(float64(product.Price) * 1.10)
    }
    return product.Price
}