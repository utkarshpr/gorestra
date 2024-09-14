package handlers

import (
	"encoding/json"
	"my-restaurant-app/internal/models"
	"my-restaurant-app/internal/services"
	"net/http"
)

type OrderHandler struct {
	orderService *services.OrderService
}

func NewOrderHandler(orderService *services.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		models.ManageResponseOrder(w, "Method not allowed", http.StatusMethodNotAllowed, nil)
		return
	}
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload "+err.Error(), http.StatusBadRequest)
		return
	}

	orderResponse, err :=
		h.orderService.CreateOrder(&order)

	if err != nil {
		models.ManageResponseOrder(w, err.Error(), http.StatusBadRequest, nil)
		return
	}

	models.ManageResponseOrder(w, "Order Created Successfully", http.StatusOK, orderResponse)

}
