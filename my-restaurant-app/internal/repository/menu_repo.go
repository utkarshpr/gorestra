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
