package utils

import (
	"errors"
	"fmt"
	"my-restaurant-app/internal/models"
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
