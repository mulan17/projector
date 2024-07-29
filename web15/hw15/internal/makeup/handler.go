package makeup

import (
	"encoding/json"
	"net/http"
	"github.com/rs/zerolog/log"
)

type CreateOrderReqBody struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type service interface {
	GetProducts() []Product
	PlaceOrder(productID string, quantity int) (Order, error)
	GetOrders() []Order
}

type Handler struct {
	s service
}

func NewHandler(s service) Handler {
	return Handler{s: s}
}

func (h Handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products := h.s.GetProducts()

	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		log.Debug().Err(err).Msg("Failed to encode products")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h Handler) OrderProduct(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateOrderReqBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		log.Debug().Err(err).Msg("Failed to decode JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order, err := h.s.PlaceOrder(reqBody.ProductID, reqBody.Quantity)
	if err != nil {
		log.Debug().Err(err).Msg("Failed to place order")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(order)
	if err != nil {
		log.Debug().Err(err).Msg("Failed to encode order")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h Handler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders := h.s.GetOrders()

	err := json.NewEncoder(w).Encode(orders)
	if err != nil {
		log.Debug().Err(err).Msg("Failed to encode orders")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
