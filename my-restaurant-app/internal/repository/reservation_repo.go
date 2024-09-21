package repository

import (
	"database/sql"
	"errors"
	"my-restaurant-app/internal/models"
)

type ReservationRepository struct {
	db *sql.DB
}

func NewReservationRepository(db *sql.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

func (r *ReservationRepository) CreateReservastion(reser *models.ReservationRequest) (*models.ReservationResponse, error) {

	query := `insert into reservations (user_id,date_time,number_of_people,special_requests) values(?,?,?,?)`
	_, err := r.db.Exec(query, reser.UserId, reser.DateTime, reser.NumberOfPeople, reser.SpecialRequests)

	if err != nil {
		if isUniqueConstraintViolation(err) {

			return nil, errors.New("user already exists with this email or id")
		}
		return nil, err
	}
	//var rr *models.ReservationResponse
	rr := &models.ReservationResponse{
		UserId:          reser.UserId,
		SpecialRequests: reser.SpecialRequests,
		NumberOfPeople:  reser.NumberOfPeople,
		DateTime:        reser.DateTime,
	}
	query = `select id from reservations where user_id=? and date_time=?`
	rows := r.db.QueryRow(query, reser.UserId, reser.DateTime)

	if err := rows.Scan(&rr.ReservationNo); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return rr, nil

}
