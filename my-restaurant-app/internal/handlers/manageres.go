package handlers

import (
	"encoding/json"
	"my-restaurant-app/internal/models"
	"my-restaurant-app/internal/services"
	"my-restaurant-app/internal/utils"
	"net/http"
)

type ManageReservationHandler struct {
	manageReserService *services.ReserrvationService
	jwtSecret          []byte
	useService         *services.UserService
}

func NewmanageHandler(manageReserService *services.ReserrvationService, jwtSecret []byte, userService *services.UserService) *ManageReservationHandler {
	return &ManageReservationHandler{
		manageReserService: manageReserService,
		jwtSecret:          jwtSecret,
		useService:         userService,
	}
}

func (h *ManageReservationHandler) CreateReservastion(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		models.ManageResponseReserv(w, "Method not supported ", http.StatusBadRequest, nil)
		return
	}

	var reservation *models.ReservationRequest
	if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
		models.ManageResponseReserv(w, err.Error(), http.StatusBadRequest, nil)
		return
	}

	err := utils.ValidateResrvation(reservation)
	if err != nil {
		models.ManageResponseReserv(w, err.Error(), http.StatusUnprocessableEntity, nil)
		return
	}

	reservationResponse, err := h.manageReserService.CreateReservastion(reservation)

	if err != nil {
		models.ManageResponseReserv(w, err.Error(), http.StatusBadRequest, nil)
		return
	}
	models.ManageResponseReserv(w, "Reservation created successfully :) ", http.StatusOK, reservationResponse)

}
