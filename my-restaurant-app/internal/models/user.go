// internal/models/user.go
package models

import "github.com/golang-jwt/jwt/v5"

// User represents a user in the system.
type User struct {
	ID        int    `json:"userid"`   // Unique identifier for the user
	Username  string `json:"username"` // User's username
	Email     string `json:"email"`    // User's email address
	Password  string `json:"password"` // User's password (hashed), not included in JSON responses
	Role      string `json:"role"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"` // User's role (e.g., "customer", "admin")
}
type UserResponse struct {
	ID        int    `json:"userid"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type GenericResponse struct {
	Message map[string]string `json:"text"`
	Data    interface{}       `json:"data,omitempty"` // Data can be any type
}

// LoginRequest represents the login request payload
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GenericResponseLogin struct {
	Message map[string]string `json:"text"`
	Data    interface{}       `json:"data,omitempty"`
	Token   string            `json:"token"`
}

// CustomClaims defines the structure for the JWT claims including custom fields
type CustomClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	//FullName string `json:"full_name"`
	// Embed the standard claims
	jwt.RegisteredClaims
}

type UpdateProfile struct {
	//ID       int    `json:"userid"`   // Unique identifier for the user
	Username string `json:"username"` // User's username
	Email    string `json:"email"`    // User's email address
	Password string `json:"password"` // User's password (hashed), not included in JSON responses
	//Role     string `json:"role"`     // User's role (e.g., "customer", "admin")
}
