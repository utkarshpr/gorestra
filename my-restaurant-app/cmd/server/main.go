// cmd/server/main.go
package main

import (
	"log"

	"my-restaurant-app/internal/database"
	"my-restaurant-app/internal/handlers"
	"my-restaurant-app/internal/repository"
	"my-restaurant-app/internal/services"
	"net/http"
)

func main() {
	jwtSecret := []byte("your_secret_key")
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService, jwtSecret)

	// Initialize the HTTP server
	http.HandleFunc("/api/users/register", userHandler.RegisterUserHandler)
	http.HandleFunc("/api/users/login", userHandler.LoginUserHandler)
	http.HandleFunc("/api/users/profile", userHandler.UserProfileHandler)

	menuRepo := repository.NewMenuRepository(db)
	menuService := services.NewMenuService(menuRepo)
	menuHandler := handlers.NewMenuHandler(menuService, jwtSecret, userService)
	http.HandleFunc("/api/menu", menuHandler.CreateMenu)
	// Start the server
	log.Println("Server started at http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
