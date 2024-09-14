package services

import (
	"my-restaurant-app/internal/models"
	"my-restaurant-app/internal/repository"
)

type OrderService struct {
	orderRepo *repository.OrderRepository
}

// NewMenuService creates a new MenuRepository.
func NewOrderService(orderRepo *repository.OrderRepository) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

func (s *OrderService) CreateOrder(order *models.Order) (*models.OrderResponse, error) {
	orderResponse, err :=
		s.orderRepo.CreateOrder(order)
	if err != nil {
		return nil, err
	}
	return orderResponse, nil
}
