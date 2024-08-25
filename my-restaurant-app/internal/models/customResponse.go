package models

import (
	"encoding/json"

	"net/http"
)

func ManageResponse(w http.ResponseWriter, errString string, code int, u *UserResponse) {
	response := GenericResponse{
		Message: map[string]string{"message": errString},
		Data:    u,
	}
	beautifiedJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(beautifiedJSON)
}
