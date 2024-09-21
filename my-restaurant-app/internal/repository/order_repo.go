package repository

import (
	"database/sql"
	"errors"
	"fmt"
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

func (r *OrderRepository) GetAllOrders() ([]*models.OrderResponse, error) {
	query := `SELECT o.id, o.user_id, o.total_price, o.status, 
				oi.menu_item_id, oi.quantity, oi.price, mi.name 
			  FROM orders o
			  JOIN order_items oi ON o.id = oi.order_id
			  JOIN menu_items mi ON oi.menu_item_id = mi.id`

	// Execute the query
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Map to hold orders by their ID (since one order can have multiple items)
	orderMap := make(map[string]*models.OrderResponse)

	// Iterate over the result rows
	for rows.Next() {
		var orderID, userID, menuItemID, menuItemName, status string
		var totalPrice float64
		var quantity int
		var itemPrice float64

		// Scan the data into the respective variables
		err := rows.Scan(&orderID, &userID, &totalPrice, &status, &menuItemID, &quantity, &itemPrice, &menuItemName)
		if err != nil {
			return nil, err
		}

		// If the order doesn't exist in the map yet, create it
		if _, exists := orderMap[orderID]; !exists {
			orderMap[orderID] = &models.OrderResponse{
				ID:         orderID,
				UserId:     userID,
				TotalPrice: fmt.Sprintf("%.2f", totalPrice),
				Status:     status,
				Orders:     []models.Orders{},
			}
		}

		// Add the item to the order
		orderMap[orderID].Orders = append(orderMap[orderID].Orders, models.Orders{
			MenuItemID: menuItemID,
			Quantity:   quantity,
			Price:      fmt.Sprintf("%.2f", itemPrice),
			MenuItem:   menuItemName,
		})
	}

	// Convert the map to a slice of orders
	var orders []*models.OrderResponse
	for _, order := range orderMap {
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepository) GetAllOrdersByUser(userID int) ([]*models.OrderResponse, error) {
	query := `
	SELECT o.id AS order_id, o.user_id, o.total_price, o.status, 
	       oi.menu_item_id, oi.quantity, oi.price, mi.name 
	FROM orders o 
	JOIN order_items oi ON o.id = oi.order_id 
	JOIN menu_items mi ON oi.menu_item_id = mi.id 
	WHERE o.user_id = ?
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orderMap := make(map[string]*models.OrderResponse)
	var orderList []*models.OrderResponse

	for rows.Next() {
		var orderID, menuItemID, menuItemName, status string
		var totalPrice string
		var quantity int
		var price string

		err = rows.Scan(&orderID, &userID, &totalPrice, &status, &menuItemID, &quantity, &price, &menuItemName)
		if err != nil {
			return nil, err
		}

		// Check if the order already exists in the map
		if _, exists := orderMap[orderID]; !exists {
			// If not, create a new order response
			orderMap[orderID] = &models.OrderResponse{
				ID:         orderID,
				UserId:     strconv.Itoa(userID),
				TotalPrice: totalPrice,
				Status:     status,
				Orders:     []models.Orders{},
			}
			orderList = append(orderList, orderMap[orderID])
		}

		// Create an order item
		orderItem := models.Orders{
			MenuItemID: menuItemID,
			Quantity:   quantity,
			Price:      price,
			MenuItem:   menuItemName,
		}

		// Append the item to the order
		orderMap[orderID].Orders = append(orderMap[orderID].Orders, orderItem)
	}

	return orderList, nil
}

func (r *OrderRepository) UpdateOrder(order *models.Order) ([]*models.OrderResponse, error) {

	query := "select status from orders where id=?"

	rows, err := r.db.Query(query, order.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var status string
	for rows.Next() {
		err = rows.Scan(&status)
		if err != nil {
			return nil, err
		}
		if status == "pending" {
			items := order.Items
			for _, item := range items {
				query = "update order_items set quantity=? where menu_item_id=? and order_id=?"
				row := r.db.QueryRow(query, item.Quantity, item.MenuItemID, order.ID)
				if err := row.Err(); err != nil {
					return nil, err
				}
			}
		} else {
			err := errors.New(" ; item cannot be update cause its not in pending state")
			return nil, err
		}

	}
	if status != order.Status {
		query = "update orders set status=? where id=?"
		row := r.db.QueryRow(query, order.Status, order.ID)
		if err := row.Err(); err != nil {
			return nil, err
		}
	}
	updatedOrder, err := r.GetOrderById(order.ID) // Assuming you have a method to fetch the updated order
	if err != nil {
		return nil, err
	}

	return updatedOrder, nil

}

func (r *OrderRepository) GetOrderById(orderID string) ([]*models.OrderResponse, error) {
	//update price
	err := r.calculateTotalPrice(orderID)
	if err != nil {
		return nil, err
	}

	query := `SELECT o.id AS order_id, o.user_id, o.total_price, o.status, 
	       oi.menu_item_id, oi.quantity, oi.price, mi.name 
	FROM orders o 
	JOIN order_items oi ON o.id = oi.order_id 
	JOIN menu_items mi ON oi.menu_item_id = mi.id 
	WHERE o.id = ?`
	rows, err := r.db.Query(query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orderMap := make(map[string]*models.OrderResponse)
	var orderList []*models.OrderResponse

	for rows.Next() {
		var orderID, menuItemID, menuItemName, status, userID string
		var totalPrice string
		var quantity int
		var price string

		err = rows.Scan(&orderID, &userID, &totalPrice, &status, &menuItemID, &quantity, &price, &menuItemName)
		if err != nil {
			return nil, err
		}

		// Check if the order already exists in the map
		if _, exists := orderMap[orderID]; !exists {
			// If not, create a new order response
			orderMap[orderID] = &models.OrderResponse{
				ID:         orderID,
				UserId:     userID,
				TotalPrice: totalPrice,
				Status:     status,
				Orders:     []models.Orders{},
			}
			orderList = append(orderList, orderMap[orderID])
		}

		// Create an order item
		orderItem := models.Orders{
			MenuItemID: menuItemID,
			Quantity:   quantity,
			Price:      price,
			MenuItem:   menuItemName,
		}

		// Append the item to the order
		orderMap[orderID].Orders = append(orderMap[orderID].Orders, orderItem)
	}

	return orderList, nil
}

func (r *OrderRepository) calculateTotalPrice(orderID string) error {

	query := `select quantity ,price from order_items where order_id = ?`
	rows, err := r.db.Query(query, orderID)
	if err != nil {
		return err
	}
	defer rows.Close()
	var total_price float64
	var quantity, Price float64
	for rows.Next() {

		rows.Scan(&quantity, &Price)
		total_price = total_price + (quantity * Price)
	}
	total_price = (.18 * total_price) + total_price

	query = "update orders set total_price=? where id=?"
	row := r.db.QueryRow(query, total_price, orderID)
	if err := row.Err(); err != nil {
		return err
	}
	return nil

}
