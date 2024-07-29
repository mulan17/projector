// Створити REST API на три ендпоінти на вільну тему за тришаровою архітектурою. 
// Це не має бути лише CRUD, потрібно додати певні бізнес логіки. Авторизацію користувача можна не реалізовувати.
// Приклад: API для сайту тур-агенства
// * отримання списку доступних турів (назва, опис, ціна, вид транспорту)
// * замовити тур. Крім збереження замовлення моделюємо відправку email (логуємо, що відправили)
// * список замовлених турів: для замовлених турів можна додати додаткові властивості/бізнес-логіки. 
// Наприклад стан туру в залежності від поточної дати (майбутній, скоро, триває, завершився)


package main

import (
	"net/http"

	"makeup/internal/makeup"

	"github.com/rs/zerolog/log"
)

func main() {
	mux := http.NewServeMux()

	productStorage := makeup.NewInMemStorage()
	productService := makeup.NewService(productStorage)
	productHandler := makeup.NewHandler(productService)

	mux.HandleFunc("/products", productHandler.ListProducts)
	mux.HandleFunc("/order", productHandler.OrderProduct)
	mux.HandleFunc("/orders", productHandler.ListOrders)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}
}
