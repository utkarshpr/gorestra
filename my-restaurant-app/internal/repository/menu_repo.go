package repository

import (
	"database/sql"
	"errors"

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

func (r *MenuRepository) UpdateMenu(menu *models.Menu) (*models.Menu, error) {
	query := "update menu_items set name=?, description=?, price=?, category=?, image=?  where id=?"
	row := r.db.QueryRow(query, menu.Name, menu.Description, menu.Price, menu.Category, menu.Image, menu.ID)
	if err := row.Err(); err != nil {
		return nil, err
	}
	query = "SELECT id, name, description, price, category, image FROM menu_items where id=?"
	row = r.db.QueryRow(query, menu.ID)

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

func (r *MenuRepository) DeleteMenu(id string) (bool, error) {

	query := "Delete from  menu_items where id=?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, errors.New("no rows were deleted")
	}

	return true, nil
}
