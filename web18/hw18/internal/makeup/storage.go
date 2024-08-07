package makeup

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sqlx.DB
}

func NewSQLStorage(db *sql.DB) *Storage {
	return &Storage{
		db: sqlx.NewDb(db, "postgres"),
	}
}

func (s *Storage) CreateProduct(p Product) error {
	_, err := s.db.Exec("INSERT INTO products (name, description, price) VALUES ($1, $2, $3)",
		p.Name, p.Description, p.Price)
	if err != nil {
		return fmt.Errorf("creating product: %w", err)
	}
	return err
}

func (s *Storage) GetAllProducts() ([]Product, error) {
	var products []Product
	err := s.db.Select(&products, "SELECT id, name, description, price FROM products")
	if err != nil {
		return nil, fmt.Errorf("selecting products: %w", err)
	}
	return products, nil
}

func (s *Storage) GetProductByID(id string) (Product, error) {
	var product Product
	err := s.db.Get(&product, "SELECT id, name, description, price FROM products WHERE id=$1", id)
	if err != nil {
		return Product{}, fmt.Errorf("selecting product by id: %w", err)
	}
	return product, nil
}

func (s *Storage) CreateOrder(o Order) error {
	_, err := s.db.Exec("INSERT INTO orders (product_id, quantity, total_price, created_at) VALUES ($1, $2, $3, $4)",
		o.ProductID, o.Quantity, o.TotalPrice, o.CreatedAt)
	if err != nil {
		return fmt.Errorf("creating order: %w", err)
	}
	return nil
}

func (s *Storage) GetAllOrders() ([]Order, error) {
	var orders []Order
	err := s.db.Select(&orders, "SELECT id, product_id, quantity, total_price, created_at FROM orders")
	if err != nil {
		return nil, fmt.Errorf("selecting orders: %w", err)
	}
	return orders, nil
}
