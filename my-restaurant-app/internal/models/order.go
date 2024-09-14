package models

type Order struct {
	ID         string `json:"id"`
	UserID     string `json:"userId"`
	Items      []Item `json:"items"`
	TotalPrice string `json:"totalPrice"`
	Status     string `json:"status"`
}

type Item struct {
	MenuItemID string `json:"menuItemId"`
	Quantity   int    `json:"quantity"`
}

type OrderResponse struct {
	UserId     string   `json:"userId"`
	ID         string   `json:"orderId"`
	TotalPrice string   `json:"totalPrice"`
	Orders     []Orders `json:"orders"`
	Status     string   `json:"status"`
}
type Orders struct {
	MenuItemID string `json:"menuItemId"`
	Quantity   int    `json:"quantity"`
	Price      string `json:"price"`
	MenuItem   string `json:"menuItem"`
}
