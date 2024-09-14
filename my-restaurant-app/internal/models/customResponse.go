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

func ManageResponseLoginToken(w http.ResponseWriter, errString string, code int, u *UserResponse, token string) {
	response := GenericResponseLogin{
		Message: map[string]string{"message": errString},
		Data:    u,
		Token:   token,
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

func ManageResponseMenu(w http.ResponseWriter, errString string, code int, u *Menu) {
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

func ManageResponseMenus(w http.ResponseWriter, errString string, code int, u []Menu) {
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

func ManageResponseOrder(w http.ResponseWriter, errString string, code int, u *OrderResponse) {
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
