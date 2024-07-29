package makeup

import (
	"sync"
)

type InMemStorage struct {
	productsM sync.Mutex
	products  []Product

	ordersM sync.Mutex
	orders  []Order
}

func NewInMemStorage() *InMemStorage {
	return &InMemStorage{
		products: []Product{
			NewProduct("1", "Lipstick", "Red Lipstick", 15),
			NewProduct("2", "Foundation", "Liquid Foundation", 20),
			NewProduct("3", "Mascara", "Black Mascara", 10),
		},
	}
}

func (s *InMemStorage) GetAllProducts() []Product {
	s.productsM.Lock()
	defer s.productsM.Unlock()

	return s.products
}

func (s *InMemStorage) CreateOrder(o Order) {
	s.ordersM.Lock()
	defer s.ordersM.Unlock()

	s.orders = append(s.orders, o)
}

func (s *InMemStorage) GetAllOrders() []Order {
	s.ordersM.Lock()
	defer s.ordersM.Unlock()

	return s.orders
}

func (s *InMemStorage) GetProductByID(id string) (Product, bool) {
	s.productsM.Lock()
	defer s.productsM.Unlock()

	for _, p := range s.products {
		if p.ID == id {
			return p, true
		}
	}

	return Product{}, false
}
