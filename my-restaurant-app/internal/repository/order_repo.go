package repository

import (
	"database/sql"
	"my-restaurant-app/internal/models"
	"strconv"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order *models.Order) (*models.OrderResponse, error) {
	query := "INSERT INTO orders (id,user_id,total_price,status) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(query, order.ID, order.UserID, order.TotalPrice, order.Status)
	if err != nil {
		return nil, err
	}

	var total_price float64
	items := order.Items
	var menuItem models.Menu
	orderResponse := &models.OrderResponse{
		UserId: order.UserID,
		ID:     order.ID,
		Status: order.Status,
	}
	for _, item := range items {

		//fetch menu stuff
		query = "SELECT id, name, description, price, category, image FROM menu_items where id=?"
		row := r.db.QueryRow(query, item.MenuItemID)

		err = row.Scan(&menuItem.ID, &menuItem.Name, &menuItem.Description, &menuItem.Price, &menuItem.Category, &menuItem.Image)

		if err != nil {
			return nil, err
		}

		//insert the orde_item
		query := "INSERT INTO order_items (id,order_id,menu_item_id,quantity,price) VALUES (?, ?, ?, ?,?)"
		var ordItemPK = order.ID + "_" + item.MenuItemID + "_" + order.UserID
		_, err = r.db.Exec(query, ordItemPK, order.ID, item.MenuItemID, item.Quantity, menuItem.Price)
		if err != nil {
			return nil, err
		}

		// calculate totalPrice
		total_price += float64(item.Quantity) * menuItem.Price

		orderItem := models.Orders{
			MenuItemID: menuItem.ID,
			Quantity:   item.Quantity,
			Price:      strconv.FormatFloat(menuItem.Price, 'f', 2, 64),
			MenuItem:   menuItem.Name,
		}

		orderResponse.Orders = append(orderResponse.Orders, orderItem)

	}
	// Add tax and finalize total price
	total_price = (.18 * total_price) + total_price
	orderResponse.TotalPrice = strconv.FormatFloat(total_price, 'f', 2, 64)

	query = "update orders set total_price=? where id=?"
	row := r.db.QueryRow(query, total_price, order.ID)
	if err := row.Err(); err != nil {
		return nil, err
	}

	return orderResponse, nil
}
