package services

import (
	"errors"
	"fmt"

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

	if user.Username == "" || user.Email == "" || user.Password == "" || user.FirstName == "" || user.LastName == "" {
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

// GetUserProfile retrieves a user profile by username
func (s *UserService) GetUserProfile(username string) (*models.User, error) {
	user, err := s.userRepo.GetUser(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUserProfile(updateProfile *models.UpdateProfile) (*models.UserResponse, error) {

	fmt.Println(updateProfile.Email + " " + updateProfile.Password + " " + updateProfile.Username)
	user, err := s.userRepo.UpdateUserPofile(updateProfile)
	if err != nil {
		return nil, err
	}
	return user, nil

}
