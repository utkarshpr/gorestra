// cmd/server/main.go
package main

import (
	"log"

	"my-restaurant-app/internal/handlers"
	"net/http"
)

func main() {

	// Initialize the HTTP server
	http.HandleFunc("/api/users/register", handlers.RegisterUser)
	//http.HandleFunc("/api/users/login", handlers.LoginUser)

	// Protected routes (middleware for authentication)
	// http.HandleFunc("/api/users/profile", middleware.AuthMiddleware(handlers.GetUserProfile))
	// http.HandleFunc("/api/users/profile/update", middleware.AuthMiddleware(handlers.UpdateUserProfile))

	// You can add more routes here for menus, orders, etc.

	// Start the server
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
