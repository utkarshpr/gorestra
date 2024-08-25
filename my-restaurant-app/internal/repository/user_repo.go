// internal/repository/user_repo.go
package repository

import (
	"database/sql"
	"errors"
	"my-restaurant-app/internal/models"

	"github.com/go-sql-driver/mysql"
)

// DB is a global variable for simplicity, but you should use dependency injection in production code.
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser inserts a new user into the database.
func (r *UserRepository) CreateUser(user *models.User) error {

	query := "INSERT INTO users (username, email, password, role,user_id) VALUES (?, ?, ?, ?,?)"
	_, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.Role, user.ID)

	if err != nil {
		if isUniqueConstraintViolation(err) {

			return errors.New("user already exists with this email or id")
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
