package makeup

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
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
        log.Error().Err(err).Msg("Failed to insert product")
    }
    return err
}

func (s *Storage) GetAllProducts() []Product {
	var products []Product
	s.db.Select(&products, "SELECT id, name, description, price FROM products")
	return products
}

func (s *Storage) GetProductByID(id string) (Product, bool) {
	var product Product
	err := s.db.Get(&product, "SELECT id, name, description, price FROM products WHERE id=$1", id)
	return product, err == nil
}

func (s *Storage) CreateOrder(o Order) {
	s.db.Exec("INSERT INTO orders (product_id, quantity, total_price, created_at) VALUES ($1, $2, $3, $4)",
		o.ProductID, o.Quantity, o.TotalPrice, o.CreatedAt)
}

func (s *Storage) GetAllOrders() []Order {
	var orders []Order
	s.db.Select(&orders, "SELECT id, product_id, quantity, total_price, created_at FROM orders")
	return orders
}

