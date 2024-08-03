package main

import (
	"database/sql"
	"net/http"
	"os"

	"makeup/internal/makeup"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {
	mux := http.NewServeMux()

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://admin:anastasiya@db:5432/makeupdb"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database")
	}

	productStorage := makeup.NewSQLStorage(db)
	productService := makeup.NewService(productStorage)
	productHandler := makeup.NewHandler(productService)

	mux.HandleFunc("/products", productHandler.CreateProduct)
	mux.HandleFunc("/products/list", productHandler.ListProducts)
	mux.HandleFunc("/order", productHandler.OrderProduct)
	mux.HandleFunc("/orders", productHandler.ListOrders)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}
}
