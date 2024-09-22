package handlers

import (
	"encoding/json"
	"my-restaurant-app/internal/models"
	"my-restaurant-app/internal/services"
	"my-restaurant-app/internal/utils"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
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

func (h *ManageReservationHandler) GetAllReservations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		models.ManageResponseReserv(w, "Method not allowed ", http.StatusMethodNotAllowed, nil)
		return
	}
	role := h.Authorization(w, r)
	if role == "admin" {
		rr, err := h.manageReserService.GetAllReservations()
		if err != nil {
			models.ManageResponseReserv(w, err.Error(), http.StatusBadRequest, nil)
			return
		}
		models.ManageResponseReservAll(w, "Reservation created successfully :) ", http.StatusOK, rr)
	} else {
		models.ManageResponseReserv(w, "Only Admin can get all reservation", http.StatusBadRequest, nil)
	}
}

func (h *ManageReservationHandler) UpdateRemoveGetReservationByID(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		h.GetReservationById(w, r)
	} else if r.Method == http.MethodPut {
		h.UpdateReservationByID(w, r)
	} else if r.Method == http.MethodDelete {
		h.DeletedReservationByID(w, r)
	} else {
		models.ManageResponseReserv(w, "Method not allowed ", http.StatusMethodNotAllowed, nil)
	}

}

func (h *ManageReservationHandler) GetReservationById(w http.ResponseWriter, r *http.Request) {

	userID := r.URL.Query().Get("userID")
	if len(userID) == 0 {
		models.ManageResponseReserv(w, "Error: userID is null", http.StatusBadRequest, nil)
		return
	}
	rr, err := h.manageReserService.GetAllReservationsById(userID)

	if err != nil {
		models.ManageResponseReserv(w, "Error: "+err.Error(), http.StatusBadRequest, nil)
		return
	}
	models.ManageResponseReserv(w, "Reservation from ID fetched successfully ", http.StatusOK, rr)
}

func (h *ManageReservationHandler) UpdateReservationByID(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")
	role := h.Authorization(w, r)
	if role == "admin" {

		if len(userID) == 0 {
			models.ManageResponseReserv(w, "Error: userID is null", http.StatusBadRequest, nil)
			return
		}

		var reservation *models.UpdateReservationRequest
		if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
			models.ManageResponseReserv(w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		err := utils.ValidateResrvationUpdate(reservation)
		if err != nil {
			models.ManageResponseReserv(w, err.Error(), http.StatusUnprocessableEntity, nil)
			return
		}
		rr, err := h.manageReserService.UpdateReservationByID(userID, reservation)

		if err != nil {
			models.ManageResponseReserv(w, "Error: "+err.Error(), http.StatusBadRequest, nil)
			return
		}
		models.ManageResponseReserv(w, "Reservation from ID Updated successfully ", http.StatusOK, rr)
	} else {
		models.ManageResponseReserv(w, "Only Admin can update the reservation", http.StatusBadRequest, nil)
	}

}

func (h *ManageReservationHandler) DeletedReservationByID(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")
	DateTime := r.URL.Query().Get("DateTime")
	if len(userID) == 0 || len(DateTime) == 0 {
		models.ManageResponseReserv(w, "Error: userID or Date and time is null", http.StatusBadRequest, nil)
		return
	}
	role := h.Authorization(w, r)
	if role == "admin" {
		err := h.manageReserService.DeletedReservationByID(userID, DateTime)

		if err != nil {
			models.ManageResponseReserv(w, "Error: "+err.Error(), http.StatusBadRequest, nil)
			return
		}
		if DateTime == "0" {
			models.ManageResponseReserv(w, "All Reservation from ID "+userID+" deleted successfully ", http.StatusOK, nil)
		} else {
			models.ManageResponseReserv(w, "Reservation from ID "+userID+" for date "+DateTime+" deleted successfully ", http.StatusOK, nil)
		}
	} else {
		models.ManageResponseReserv(w, "Only Admin can delete the reservation", http.StatusBadRequest, nil)
	}

}
func (h *ManageReservationHandler) Authorization(w http.ResponseWriter, r *http.Request) string {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		models.ManageResponseMenu(w, "Authorization header missing", http.StatusUnauthorized, nil)
		return ""
	}

	// Extract token from header
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		models.ManageResponseMenu(w, "Bearer token missing", http.StatusUnauthorized, nil)
		return ""
	}

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrNoLocation
		}
		return h.jwtSecret, nil
	})
	if err != nil || !token.Valid {
		models.ManageResponseMenu(w, "Invalid token", http.StatusUnauthorized, nil)
		return ""
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		models.ManageResponseMenu(w, "Invalid token claims", http.StatusUnauthorized, nil)
		return ""
	}

	username, ok := claims["sub"].(string)

	if !ok {
		models.ManageResponseMenu(w, "Username not found in token ", http.StatusUnauthorized, nil)
		return ""
	}

	// Get user profile
	_, err = h.useService.GetUserProfile(username)
	if err != nil {
		models.ManageResponseMenu(w, "profile no exist please register"+err.Error(), http.StatusNotFound, nil)
		return ""
	}
	role, ok := claims["role"].(string)

	if !ok {
		models.ManageResponseMenu(w, "Username not found in token ", http.StatusUnauthorized, nil)
		return ""
	}

	return role
}
