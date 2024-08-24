// internal/handlers/user.go
package handlers

import (
	"encoding/json"
	"fmt"
	"my-restaurant-app/internal/auth"
	"my-restaurant-app/internal/models"
	"my-restaurant-app/internal/repository"
	"net/http"
)

// RegisterUser handles user registration.
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode the request body into the user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Hash the user's password before saving
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	// Save the user to the database
	err = repository.CreateUser(user)
	if err != nil {
		fmt.Print(err)
		if err.Error() == "user already exists with this email" {
			http.Error(w, "User already exists", http.StatusConflict)
		} else {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
		}
		return
	}

	// Respond with the created user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
