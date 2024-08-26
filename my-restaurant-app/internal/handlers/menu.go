// internal/handlers/user.go
package handlers

import (
	"io"
	"my-restaurant-app/internal/models"
	"my-restaurant-app/internal/services"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	//go"github.com/dgrijalva/jwt-go"
)

// MenuHandler handles user-related HTTP requests.
type MenuHandler struct {
	menuService *services.MenuService
	jwtSecret   []byte
	useService  *services.UserService
}

// type UserRegisterResponse struct {
// 	message      map[string]string
// 	userResponse models.UserResponse
// }

// NewMenuHandler creates a new MenuHandler.
func NewMenuHandler(menuService *services.MenuService, jwtSecret []byte, userService *services.UserService) *MenuHandler {
	return &MenuHandler{
		menuService: menuService,
		jwtSecret:   jwtSecret,
		useService:  userService,
	}
}

func (h *MenuHandler) CreateMenu(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		models.ManageResponseMenu(w, "Invalid request Method ", http.StatusBadRequest, nil)
		return
	}

	var menu *models.Menu

	// Get token from Authorization header
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		models.ManageResponseMenu(w, "Authorization header missing", http.StatusUnauthorized, nil)
		return
	}

	// Extract token from header
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		models.ManageResponseMenu(w, "Bearer token missing", http.StatusUnauthorized, nil)
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
		models.ManageResponseMenu(w, "Invalid token", http.StatusUnauthorized, nil)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		models.ManageResponseMenu(w, "Invalid token claims", http.StatusUnauthorized, nil)
		return
	}

	username, ok := claims["sub"].(string)

	if !ok {
		models.ManageResponseMenu(w, "Username not found in token ", http.StatusUnauthorized, nil)
		return
	}

	// Get user profile
	_, err = h.useService.GetUserProfile(username)
	if err != nil {
		models.ManageResponseMenu(w, "profile no exist please register"+err.Error(), http.StatusNotFound, nil)
		return
	}
	role, ok := claims["role"].(string)

	if !ok {
		models.ManageResponseMenu(w, "Username not found in token ", http.StatusUnauthorized, nil)
		return
	}

	if role != "admin" {
		models.ManageResponseMenu(w, "invalid token or not admin", http.StatusUnauthorized, nil)
		return
	}

	r.ParseMultipartForm(10 << 20) // 10 MB file size limit

	// Get file handler
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	// Create the project-specific directory if it doesn't exist
	projectDir := "uploads"
	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
		err := os.MkdirAll(projectDir, os.ModePerm)
		if err != nil {
			http.Error(w, "Unable to create directory on server", http.StatusInternalServerError)
			return
		}
	}

	// Create file on server in the project-specific directory
	filePath := filepath.Join(projectDir, handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to create the file on server", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the server's file system
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Unable to save the file", http.StatusInternalServerError)
		return
	}

	// Parse other form fields
	menuid := r.FormValue("menuid")
	name := r.FormValue("name")
	description := r.FormValue("description")
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		models.ManageResponseMenu(w, "Price conversion"+err.Error(), http.StatusBadRequest, nil)
	}
	category := r.FormValue("category")

	menu = &models.Menu{
		ID:          menuid,
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
		Image:       filePath, // Store the file path in the database
	}

	menu, err = h.menuService.CreateMenu(menu)
	if err != nil {
		models.ManageResponseMenu(w, err.Error(), http.StatusBadRequest, nil)
		return
	}

	models.ManageResponseMenu(w, "Menu Created in successfully", http.StatusCreated, menu)

}
