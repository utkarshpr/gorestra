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

func (s *ReserrvationService) GetAllReservations() ([]*models.ReservationResponse, error) {
	rr, err := s.reservationRepo.GetAllReservations()
	if err != nil {
		return nil, err
	}
	return rr, nil
}

func (s *ReserrvationService) GetAllReservationsById(userID string) (*models.ReservationResponse, error) {
	rr, err := s.reservationRepo.GetAllReservationsById(userID)
	if err != nil {
		return nil, err
	}
	return rr, nil
}

func (s *ReserrvationService) UpdateReservationByID(userID string, reservation *models.UpdateReservationRequest) (*models.ReservationResponse, error) {
	rr, err := s.reservationRepo.UpdateReservationByID(userID, reservation)
	if err != nil {
		return nil, err
	}
	return rr, nil
}

func (s *ReserrvationService) DeletedReservationByID(userID string, DateTime string) error {
	err := s.reservationRepo.DeletedReservationByID(userID, DateTime)
	if err != nil {
		return err
	}
	return nil
}
