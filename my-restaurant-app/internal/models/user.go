// internal/models/user.go
package models

// User represents a user in the system.
type User struct {
	ID       int    `json:"userid"`   // Unique identifier for the user
	Username string `json:"username"` // User's username
	Email    string `json:"email"`    // User's email address
	Password string `json:"password"` // User's password (hashed), not included in JSON responses
	Role     string `json:"role"`     // User's role (e.g., "customer", "admin")
}
type UserResponse struct {
	ID       int    `json:"userid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
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
