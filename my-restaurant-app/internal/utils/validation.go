package utils

import (
	"errors"
	"fmt"
	"my-restaurant-app/internal/models"
	"strconv"
	"time"
)

func ValidateOrder(order *models.Order) error {

	// Check if ID is empty
	if order.ID == "" {
		return errors.New("order ID cannot be empty")
	}

	// Check if UserID is empty
	if order.UserID == "" {
		return errors.New("user ID cannot be empty")
	}

	// Check if TotalPrice is empty
	// if order.TotalPrice == "" {
	// 	return errors.New("total price cannot be empty")
	// }

	// Check if Status is valid (must be one of the predefined statuses)
	validStatuses := map[string]bool{
		"pending":   true,
		"preparing": true,
		"completed": true,
		"canceled":  true,
	}

	if _, ok := validStatuses[order.Status]; !ok {
		return fmt.Errorf("invalid status: %s", order.Status)
	}

	// Check if Items slice is empty
	if len(order.Items) == 0 {
		return errors.New("order must contain at least one item")
	}

	// Validate each item in the order
	for _, item := range order.Items {
		if item.MenuItemID == "" {
			return errors.New("menu item ID cannot be empty")
		}
		if item.Quantity <= 0 {
			return errors.New("quantity must be greater than 0")
		}
	}

	// If all validations pass, return nil (no error)
	return nil

}

func ValidateResrvation(r *models.ReservationRequest) error {
	// Validate UserId
	if r.UserId == "" {
		return errors.New("userId cannot be empty")
	}

	// Validate DateTime - check if it's in the correct format and a future date
	const layout = "2006-01-02 15:04:05"
	_, err := time.Parse(layout, r.DateTime)
	if err != nil {
		return fmt.Errorf("dateTime must be in the format YYYY-MM-DD HH:MM:SS")
	}

	// Validate NumberOfPeople - check if it's a valid number and greater than 0
	numPeople, err := strconv.Atoi(r.NumberOfPeople)
	if err != nil || numPeople <= 0 {
		return errors.New("numberOfPeople must be a valid positive integer")
	}

	// Validate SpecialRequests - optional field, but we can add a length limit
	if len(r.SpecialRequests) > 200 {
		return errors.New("specialRequests cannot exceed 200 characters")
	}

	// Optionally, add more validations (e.g., check if UserId exists in DB)

	return nil

}

func ValidateResrvationUpdate(r *models.UpdateReservationRequest) error {
	// Validate UserId

	// Validate DateTime - check if it's in the correct format and a future date
	const layout = "2006-01-02 15:04:05"
	_, err := time.Parse(layout, r.DateTime)
	if err != nil {
		return fmt.Errorf("dateTime must be in the format YYYY-MM-DD HH:MM:SS")
	}

	// Validate NumberOfPeople - check if it's a valid number and greater than 0
	numPeople, err := strconv.Atoi(r.NumberOfPeople)
	if err != nil || numPeople <= 0 {
		return errors.New("numberOfPeople must be a valid positive integer")
	}

	// Validate SpecialRequests - optional field, but we can add a length limit
	if len(r.SpecialRequests) > 200 {
		return errors.New("specialRequests cannot exceed 200 characters")
	}

	// Optionally, add more validations (e.g., check if UserId exists in DB)

	return nil

}
