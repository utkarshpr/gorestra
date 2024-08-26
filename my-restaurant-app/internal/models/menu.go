package models

type Menu struct {
	ID          string  `json:"menuid"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
}
