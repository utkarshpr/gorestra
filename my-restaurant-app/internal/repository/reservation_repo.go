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

	query := `select count(*) from reservations where user_id=? and date_time= ?`
	var count int
	err1 := r.db.QueryRow(query, reser.UserId, reser.DateTime).Scan(&count)
	if err1 != nil {
		return nil, err1
	}
	if count > 0 {
		err := errors.New("user and date already exist in reservation chart")
		return nil, err
	}

	query = `insert into reservations (user_id,date_time,number_of_people,special_requests) values(?,?,?,?)`
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
	query = `select r.id,u.first_name,u.last_name from reservations r join users u on u.user_id=r.user_id where r.user_id=? and date_time=?`
	rows := r.db.QueryRow(query, reser.UserId, reser.DateTime)

	if err := rows.Scan(&rr.ReservationNo, &rr.FirstName, &rr.LastName); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return rr, nil

}

func (r *ReservationRepository) GetAllReservations() ([]*models.ReservationResponse, error) {
	query := `select id,u.user_id,date_time,number_of_people ,special_requests,first_name,last_name 
				from reservations r join users u on u.user_id=r.user_id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rr []*models.ReservationResponse

	for rows.Next() {
		var ReservationNo, UserId, DateTime, NumberOfPeople, SpecialRequests, FirstName, LastName string
		err := rows.Scan(&ReservationNo, &UserId, &DateTime, &NumberOfPeople, &SpecialRequests, &FirstName, &LastName)
		if err != nil {
			return nil, err
		}

		//insert into map via user_id

		rowData := &models.ReservationResponse{
			ReservationNo:   ReservationNo,
			UserId:          UserId,
			DateTime:        DateTime,
			NumberOfPeople:  NumberOfPeople,
			SpecialRequests: SpecialRequests,
			FirstName:       FirstName,
			LastName:        LastName,
		}
		rr = append(rr, rowData)

	}

	return rr, nil
}

func (r *ReservationRepository) GetAllReservationsById(userID string) (*models.ReservationResponse, error) {
	query := `select id,u.user_id,date_time,number_of_people ,special_requests,first_name,last_name 
	from reservations r join users u on u.user_id=r.user_id and u.user_id =?`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// if !rows.Next() {
	// 	err := errors.New("no reservation with this id is present")
	// 	return nil, err
	// }

	var rr *models.ReservationResponse

	for rows.Next() {
		var ReservationNo, UserId, DateTime, NumberOfPeople, SpecialRequests, FirstName, LastName string
		err := rows.Scan(&ReservationNo, &UserId, &DateTime, &NumberOfPeople, &SpecialRequests, &FirstName, &LastName)
		if err != nil {
			return nil, err
		}

		//insert into map via user_id

		rr = &models.ReservationResponse{
			ReservationNo:   ReservationNo,
			UserId:          UserId,
			DateTime:        DateTime,
			NumberOfPeople:  NumberOfPeople,
			SpecialRequests: SpecialRequests,
			FirstName:       FirstName,
			LastName:        LastName,
		}

	}
	if rr == nil {
		err := errors.New("no reservation with this id is present")
		return nil, err
	}

	return rr, nil
}

func (r *ReservationRepository) UpdateReservationByID(userID string, reservation *models.UpdateReservationRequest) (*models.ReservationResponse, error) {
	query := `update reservations set date_time=?, special_requests=?, number_of_people=? where user_id=?`

	_, err := r.db.Query(query, reservation.DateTime, reservation.SpecialRequests, reservation.NumberOfPeople, userID)
	if err != nil {
		return nil, err
	}

	query = `select id,u.user_id,date_time,number_of_people ,special_requests,first_name,last_name 
	from reservations r join users u on u.user_id=r.user_id and u.user_id =?`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rr *models.ReservationResponse

	for rows.Next() {
		var ReservationNo, UserId, DateTime, NumberOfPeople, SpecialRequests, FirstName, LastName string
		err := rows.Scan(&ReservationNo, &UserId, &DateTime, &NumberOfPeople, &SpecialRequests, &FirstName, &LastName)
		if err != nil {
			return nil, err
		}

		//insert into map via user_id

		rr = &models.ReservationResponse{
			ReservationNo:   ReservationNo,
			UserId:          UserId,
			DateTime:        DateTime,
			NumberOfPeople:  NumberOfPeople,
			SpecialRequests: SpecialRequests,
			FirstName:       FirstName,
			LastName:        LastName,
		}

	}
	if rr == nil {
		err := errors.New("no reservation with this id is present")
		return nil, err
	}

	return rr, nil
}

func (r *ReservationRepository) DeletedReservationByID(userID string, dateTime string) error {

	var query string
	var err error
	var result sql.Result
	if dateTime != "0" {
		query = `delete from reservations where user_id=? and date_time=?`
		result, err = r.db.Exec(query, userID, dateTime)
	} else {
		query = `delete from reservations where user_id=? `
		result, err = r.db.Exec(query, userID)
	}
	if err != nil {
		return err
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no reservation found for user_id at date_time")
	}

	return nil
}
