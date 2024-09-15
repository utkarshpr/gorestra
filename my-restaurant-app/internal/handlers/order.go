package handlers

import (
	"encoding/json"
	"my-restaurant-app/internal/models"
	"my-restaurant-app/internal/services"
	"my-restaurant-app/internal/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type OrderHandler struct {
	orderService *services.OrderService
	jwtSecret    []byte
	useService   *services.UserService
}

func NewOrderHandler(orderService *services.OrderService, jwtSecret []byte, userService *services.UserService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
		jwtSecret:    jwtSecret,
		useService:   userService,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		models.ManageResponseOrder(w, "Method not allowed", http.StatusMethodNotAllowed, nil)
		return
	}
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		models.ManageResponseOrder(w, "Invalid request payload "+err.Error(), http.StatusBadRequest, nil)
		return
	}

	//validation
	err := utils.ValidateOrder(&order)
	if err != nil {
		models.ManageResponseOrder(w, err.Error(), http.StatusBadRequest, nil)
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

func (h *OrderHandler) FetchAllOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		models.ManageResponseOrder(w, "Invalid request Method ", http.StatusMethodNotAllowed, nil)
		return
	}

	role := h.Authorization(w, r)
	if role == "admin" {
		orderResponse, err := h.orderService.GetAllOrders()

		if err != nil {
			models.ManageResponseOrder(w, err.Error(), http.StatusBadRequest, nil)
			return
		}
		models.ManageResponseOrders(w, "All Orders fetch successfully", http.StatusOK, orderResponse)

	} else {
		models.ManageResponseOrder(w, "All orders can only be fetch by admin ", http.StatusBadRequest, nil)
	}

}

func (h *OrderHandler) GetOrdersByUserID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		models.ManageResponseMenu(w, "Invalid Request Method", http.StatusMethodNotAllowed, nil)
		return
	}
	role := h.Authorization(w, r)

	if role == "admin" {
		//var menus *models.Menu
		userID := r.URL.Query().Get("userID")
		if userID == "" {
			models.ManageResponseMenu(w, "ID is required", http.StatusBadRequest, nil)
			return
		}
		num, err := strconv.Atoi(userID)
		if err != nil {
			models.ManageResponseMenu(w, err.Error(), http.StatusBadRequest, nil)
		}
		orders, err := h.orderService.GetOrdersByUserID(num)
		if err != nil {
			models.ManageResponseMenu(w, "Failed to fetch the menu", http.StatusBadRequest, nil)
			return
		}
		if orders == nil {
			models.ManageResponseMenu(w, "Failed to fetch the menu as id donot exists", http.StatusBadRequest, nil)
			return
		}
		models.ManageResponseOrders(w, "All orders can only be fetch by admin ", http.StatusBadRequest, orders)
	} else {
		models.ManageResponseOrder(w, "All orders By user can only be fetch by admin ", http.StatusBadRequest, nil)
	}

}

func (h *OrderHandler) Authorization(w http.ResponseWriter, r *http.Request) string {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		models.ManageResponseOrder(w, "Authorization header missing", http.StatusUnauthorized, nil)
		return ""
	}

	// Extract token from header
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		models.ManageResponseOrder(w, "Bearer token missing", http.StatusUnauthorized, nil)
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
		models.ManageResponseOrder(w, "Invalid token", http.StatusUnauthorized, nil)
		return ""
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		models.ManageResponseOrder(w, "Invalid token claims", http.StatusUnauthorized, nil)
		return ""
	}

	username, ok := claims["sub"].(string)

	if !ok {
		models.ManageResponseOrder(w, "Username not found in token ", http.StatusUnauthorized, nil)
		return ""
	}

	// Get user profile
	_, err = h.useService.GetUserProfile(username)
	if err != nil {
		models.ManageResponseOrder(w, "profile no exist please register"+err.Error(), http.StatusNotFound, nil)
		return ""
	}
	role, ok := claims["role"].(string)

	if !ok {
		models.ManageResponseOrder(w, "Username not found in token ", http.StatusUnauthorized, nil)
		return ""
	}

	return role
}
