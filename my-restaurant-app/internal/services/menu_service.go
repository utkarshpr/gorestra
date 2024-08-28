package services

import (
	"errors"

	"my-restaurant-app/internal/models"
	"my-restaurant-app/internal/repository"
)

// MenuService provides user-related services.
type MenuService struct {
	menuRepo *repository.MenuRepository
}

// NewMenuService creates a new MenuRepository.
func NewMenuService(menuRepo *repository.MenuRepository) *MenuService {
	return &MenuService{menuRepo: menuRepo}
}

// RegisterUser handles the user registration logic.
func (s *MenuService) CreateMenu(menu *models.Menu) (*models.Menu, error) {

	if menu.ID == "" || menu.Name == "" || menu.Price == 0 {
		return nil, errors.New("missing required fields")
	}

	// Save the user
	menu, err := s.menuRepo.CreateMenu(menu)
	return menu, err
}

func (s *MenuService) FetchAllMenu() ([]models.Menu, error) {
	menu, err := s.menuRepo.FetchAllMenu()
	return menu, err
}

func (s *MenuService) FetchMenu(id string) (*models.Menu, error) {
	menu, err := s.menuRepo.FetchMenu(id)
	return menu, err
}
