// internal/models/user.go
package models

// User represents a user in the system.
type User struct {
	ID       int    `json:"id"`       // Unique identifier for the user
	Username string `json:"username"` // User's username
	Email    string `json:"email"`    // User's email address
	Password string `json:"-"`        // User's password (hashed), not included in JSON responses
	Role     string `json:"role"`     // User's role (e.g., "customer", "admin")
}
