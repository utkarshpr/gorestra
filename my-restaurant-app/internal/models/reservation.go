package models

type ReservationResponse struct {
	ReservationNo   string `json:"reservationNo"`
	UserId          string `json:"userId"`
	Name            string `json:"name"`
	DateTime        string `json:"dateTime"`
	NumberOfPeople  string `json:"numberOfPeople"`
	SpecialRequests string `json:"specialRequests"`
}

type ReservationRequest struct {
	UserId          string `json:"userId"`
	DateTime        string `json:"dateTime"`
	NumberOfPeople  string `json:"numberOfPeople"`
	SpecialRequests string `json:"specialRequests"`
}
