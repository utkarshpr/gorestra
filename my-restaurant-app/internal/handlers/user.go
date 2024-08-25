// internal/handlers/user.go
package handlers

import (
	"encoding/json"
	"my-restaurant-app/internal/models"
	"my-restaurant-app/internal/services"
	"net/http"
)

// UserHandler handles user-related HTTP requests.
type UserHandler struct {
	userService *services.UserService
}

// type UserRegisterResponse struct {
// 	message      map[string]string
// 	userResponse models.UserResponse
// }

// NewUserHandler creates a new UserHandler.
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// RegisterUserHandler handles the user registration HTTP request.
func (h *UserHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.userService.RegisterUser(&user); err != nil {

		models.ManageResponse(w, err.Error(), http.StatusBadRequest, nil)

		return
	}
	userResponse := models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}

	models.ManageResponse(w, "User registered successfully", http.StatusCreated, &userResponse)
}

func (h *UserHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var login models.LoginRequest

	decoder := json.NewDecoder(r.Body)
	// Make decoder strict
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&login); err != nil {

		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	user, err := h.userService.LoginUser(&login)
	if err != nil {

		models.ManageResponse(w, err.Error(), http.StatusBadRequest, nil)
		return
	}

	models.ManageResponse(w, "User Logged in successfully", http.StatusCreated, user)

}
