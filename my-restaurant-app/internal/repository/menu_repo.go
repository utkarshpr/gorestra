package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"my-restaurant-app/internal/models"
)

type MenuRepository struct {
	db *sql.DB
}

func NewMenuRepository(db *sql.DB) *MenuRepository {
	return &MenuRepository{db: db}
}

func (r *MenuRepository) CreateMenu(menu *models.Menu) (*models.Menu, error) {

	query := "INSERT INTO menu_items (id, name, description, price,category,image) VALUES (?, ?, ?, ?,?,?)"
	_, err := r.db.Exec(query, menu.ID, menu.Name, menu.Description, menu.Price, menu.Category, menu.Image)

	if err != nil {
		if isUniqueConstraintViolation(err) {

			return nil, errors.New("Menu already exists with this id")
		}
		return nil, err
	}

	return menu, nil
}
func (r *MenuRepository) FetchAllMenu() ([]models.Menu, error) {
	query := "SELECT id, name, description, price, category, image FROM menu_items"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menuItems []models.Menu

	for rows.Next() {
		var menuItem models.Menu
		err := rows.Scan(&menuItem.ID, &menuItem.Name, &menuItem.Description, &menuItem.Price, &menuItem.Category, &menuItem.Image)
		if err != nil {
			return nil, err
		}
		menuItems = append(menuItems, menuItem)
	}
	fmt.Println(menuItems)

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return menuItems, nil
}

func (r *MenuRepository) FetchMenu(id string) (*models.Menu, error) {
	query := "SELECT id, name, description, price, category, image FROM menu_items where id=?"
	row := r.db.QueryRow(query, id)

	var menuItem models.Menu
	err := row.Scan(&menuItem.ID, &menuItem.Name, &menuItem.Description, &menuItem.Price, &menuItem.Category, &menuItem.Image)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No result found
		}
		return nil, err
	}

	return &menuItem, nil
}
