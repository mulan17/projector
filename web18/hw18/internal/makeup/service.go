package makeup

import (
    "errors"
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
    return s.s.CreateProduct(p)
}

func (s *Service) GetAllProducts() ([]Product, error) {
    return s.s.GetAllProducts()
}

func (s *Service) GetProductByID(id string) (Product, error) {
    return s.s.GetProductByID(id)
}

func (s *Service) CreateOrder(o Order) error {
    return s.s.CreateOrder(o)
}

func (s *Service) GetAllOrders() ([]Order, error) {
    return s.s.GetAllOrders()
}

func (s *Service) PlaceOrder(productID string, quantity int) (Order, error) {
    product, err := s.s.GetProductByID(productID)
    if err != nil {
        return Order{}, errors.New("product not found")
    }

    totalPrice := s.getPrice(product) * quantity
    order := NewOrder(productID, quantity, totalPrice)
    s.s.CreateOrder(order)

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