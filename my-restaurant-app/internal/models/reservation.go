package models

type ReservationResponse struct {
	ReservationNo   string `json:"reservationNo"`
	UserId          string `json:"userId"`
	DateTime        string `json:"dateTime"`
	NumberOfPeople  string `json:"numberOfPeople"`
	SpecialRequests string `json:"specialRequests"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
}

type ReservationRequest struct {
	UserId          string `json:"userId"`
	DateTime        string `json:"dateTime"`
	NumberOfPeople  string `json:"numberOfPeople"`
	SpecialRequests string `json:"specialRequests"`
}

type UpdateReservationRequest struct {
	DateTime        string `json:"dateTime"`
	NumberOfPeople  string `json:"numberOfPeople"`
	SpecialRequests string `json:"specialRequests"`
}
