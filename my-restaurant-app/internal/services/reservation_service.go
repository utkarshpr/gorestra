package services

import (
	"my-restaurant-app/internal/models"
	"my-restaurant-app/internal/repository"
)

type ReserrvationService struct {
	reservationRepo *repository.ReservationRepository
}

// NewUserService creates a new UserService.
func NewReserrvationService(reservationRepo *repository.ReservationRepository) *ReserrvationService {
	return &ReserrvationService{reservationRepo: reservationRepo}
}

func (s *ReserrvationService) CreateReservastion(reser *models.ReservationRequest) (*models.ReservationResponse, error) {

	reservationResponse, err := s.reservationRepo.CreateReservastion(reser)
	if err != nil {
		return nil, err
	}
	return reservationResponse, nil
}
