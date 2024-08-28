// internal/handlers/user.go
package handlers

import (
	"fmt"
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

	role := h.Authorization(w, r)

	if role == "admin" {
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

}

func (h *MenuHandler) FetchAllMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		models.ManageResponseMenu(w, "Invalid Request Method", http.StatusMethodNotAllowed, nil)
		return
	}

	role := h.Authorization(w, r)

	if role == "admin" || role == "customer" {

		var menus []models.Menu
		menus, err := h.menuService.FetchAllMenu()
		if err != nil {
			models.ManageResponseMenu(w, "Failed to fetch the menu", http.StatusBadRequest, nil)
			return
		}
		models.ManageResponseMenus(w, "Fetch menu successfully", http.StatusOK, menus)

	}
}
func (h *MenuHandler) FetchMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		models.ManageResponseMenu(w, "Invalid Request Method", http.StatusMethodNotAllowed, nil)
		return
	}

	var menus *models.Menu
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		models.ManageResponseMenu(w, "ID is required", http.StatusBadRequest, nil)
		return
	}

	menus, err := h.menuService.FetchMenu(idStr)
	if err != nil {
		models.ManageResponseMenu(w, "Failed to fetch the menu", http.StatusBadRequest, nil)
		return
	}
	fmt.Println(menus)
	if menus == nil {
		models.ManageResponseMenu(w, "Failed to fetch the menu as id donot exists", http.StatusBadRequest, nil)
		return
	}
	models.ManageResponseMenu(w, " Menu fetch successfully", http.StatusOK, menus)

}

func (h *MenuHandler) UpdateMenu(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		models.ManageResponseMenu(w, "Invalid request Method ", http.StatusBadRequest, nil)
		return
	}

	var menu *models.Menu

	role := h.Authorization(w, r)

	if role == "admin" {
		r.ParseMultipartForm(10 << 20) // 10 MB file size limit

		// Get file handler
		file, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Error retrieving the file"+err.Error(), http.StatusBadRequest)
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

		menu, err = h.menuService.UpdateMenu(menu)
		if err != nil {
			models.ManageResponseMenu(w, err.Error(), http.StatusBadRequest, nil)
			return
		}
		if menu == nil {
			models.ManageResponseMenu(w, "Failed to Update the menu as id donot exists", http.StatusBadRequest, nil)
			return
		}

		models.ManageResponseMenu(w, "Menu Updated in successfully", http.StatusCreated, menu)
	} else {
		models.ManageResponseMenu(w, "Updtae can only be perform by Admin", http.StatusUnauthorized, nil)
	}

}

func (h *MenuHandler) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		models.ManageResponseMenu(w, "Invalid Request Method", http.StatusMethodNotAllowed, nil)
		return
	}

	role := h.Authorization(w, r)
	if role == "admin" {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			models.ManageResponseMenu(w, "ID is required", http.StatusBadRequest, nil)
			return
		}

		done, err := h.menuService.DeleteMenu(idStr)
		if err != nil {
			models.ManageResponseMenu(w, "Failed to delete the menu "+err.Error(), http.StatusBadRequest, nil)
			return
		}

		if !done {
			models.ManageResponseMenu(w, "Failed to delete the menu as id donot exists", http.StatusBadRequest, nil)
			return
		}
		models.ManageResponseMenu(w, " Menu deleted successfully ", http.StatusOK, nil)
	} else {
		models.ManageResponseMenu(w, "Deletion can only be perform by Admin", http.StatusUnauthorized, nil)
	}

}

func (h *MenuHandler) Authorization(w http.ResponseWriter, r *http.Request) string {
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
