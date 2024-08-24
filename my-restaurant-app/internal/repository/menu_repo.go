// internal/repository/user_repo.go
package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"my-restaurant-app/internal/database"
	"my-restaurant-app/internal/models"

	"github.com/go-sql-driver/mysql"
)

// DB is a global variable for simplicity, but you should use dependency injection in production code.
var DB *sql.DB

// CreateUser adds a new user to the SQL database.
func CreateUser(user models.User) error {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()
	// Prepare the SQL statement for inserting a new user
	query := `INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?)`
	fmt.Print(query)
	// Execute the SQL query with the user's data
	_, err = db.Exec(query, user.Username, user.Email, user.Password, user.Role)
	fmt.Print(err)
	if err != nil {
		if isUniqueConstraintViolation(err) {
			return errors.New("user already exists with this email")
		}
		return err
	}

	return nil
}

// isUniqueConstraintViolation checks if an error is a unique constraint violation (specific to the database used).
func isUniqueConstraintViolation(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if mysqlErr.Number == 1062 {
			return true
		}
	}
	return false // Replace with actual implementation
}
