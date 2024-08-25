package services

import (
	"errors"

	"my-restaurant-app/internal/auth"
	"my-restaurant-app/internal/models"
	"my-restaurant-app/internal/repository"
)

// UserService provides user-related services.
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// RegisterUser handles the user registration logic.
func (s *UserService) RegisterUser(user *models.User) error {

	if user.Username == "" || user.Email == "" || user.Password == "" {
		return errors.New("missing required fields")
	}

	// Hash the password
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Save the user
	return s.userRepo.CreateUser(user)
}

func (s *UserService) LoginUser(user *models.LoginRequest) (*models.UserResponse, error) {

	if user.Email == "" || user.Password == "" {
		return nil, errors.New("missing required fields")
	}

	// fetch the user
	userFetched, err := s.userRepo.LoginUser(user)
	if err != nil {
		return nil, err
	}

	return userFetched, nil
}
