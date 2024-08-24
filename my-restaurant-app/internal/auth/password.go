// internal/auth/password.go
package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt.
func HashPassword(password string) (string, error) {
	// GenerateFromPassword generates a bcrypt hash of the password using the default cost (bcrypt.DefaultCost)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
