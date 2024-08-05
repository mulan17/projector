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

type CreateProductReqBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type service interface {
    CreateProduct(p Product) error
	GetAllProducts() ([]Product, error)
    GetProductByID(id string) (Product, error)
    CreateOrder(o Order) error
    GetAllOrders() ([]Order, error)
	PlaceOrder(productID string, quantity int) (Order, error)
}

type Handler struct {
	s service
}

func NewHandler(s service) Handler {
	return Handler{s: s}
}

func (h Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateProductReqBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		log.Debug().Err(err).Msg("Failed to decode JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product := Product{
		Name:        reqBody.Name,
		Description: reqBody.Description,
		Price:       reqBody.Price,
	}

	err = h.s.CreateProduct(product)
	if err != nil {
		log.Debug().Err(err).Msg("Failed to create product")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h Handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.s.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
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
	orders, err := h.s.GetAllOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
