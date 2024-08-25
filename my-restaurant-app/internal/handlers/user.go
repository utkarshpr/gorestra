// internal/handlers/user.go
package handlers

import (
	"encoding/json"
	"my-restaurant-app/internal/models"
	"my-restaurant-app/internal/services"
	"net/http"
	"strings"
	"time"

	//go"github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt/v5"
)

// UserHandler handles user-related HTTP requests.
type UserHandler struct {
	userService *services.UserService
	jwtSecret   []byte
}

// type UserRegisterResponse struct {
// 	message      map[string]string
// 	userResponse models.UserResponse
// }

// NewUserHandler creates a new UserHandler.
func NewUserHandler(userService *services.UserService, jwtSecret []byte) *UserHandler {
	return &UserHandler{
		userService: userService,
		jwtSecret:   jwtSecret,
	}
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

	// Generate JWT token
	tokenString, err := h.generateJWTToken(user.Email)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	models.ManageResponseLoginToken(w, "User Logged in successfully", http.StatusCreated, user, tokenString)

}

// GetUserProfileHandler retrieves the profile of the logged-in user
func (h *UserHandler) GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		//custome response
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get token from Authorization header
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		models.ManageResponse(w, "Authorization header missing", http.StatusUnauthorized, nil)
		return
	}

	// Extract token from header
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		models.ManageResponse(w, "Bearer token missing", http.StatusUnauthorized, nil)
		return
	}

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrNoLocation
		}
		return h.jwtSecret, nil
	})
	if err != nil || !token.Valid {
		models.ManageResponse(w, "Invalid token", http.StatusUnauthorized, nil)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		models.ManageResponse(w, "Invalid token claims", http.StatusUnauthorized, nil)
		return
	}

	username, ok := claims["sub"].(string)

	if !ok {
		models.ManageResponse(w, "Username not found in token ", http.StatusUnauthorized, nil)
		return
	}

	// Get user profile
	user, err := h.userService.GetUserProfile(username)
	if err != nil {
		models.ManageResponse(w, err.Error(), http.StatusNotFound, nil)
		return
	}
	userResponse := models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}

	models.ManageResponse(w, "User profile retrieved successfully", http.StatusOK, &userResponse)
}

// Helper function to generate JWT token
func (h *UserHandler) generateJWTToken(username string) (string, error) {
	// Define token expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims, which includes the username and expiry time
	claims := jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(h.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
