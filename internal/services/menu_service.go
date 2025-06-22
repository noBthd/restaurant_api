package services

import (
	"database/sql"

	"github.com/noBthd/restaurant_api.git/internal/db"
	"github.com/noBthd/restaurant_api.git/internal/models"
)

func GetAllMenuItems()([]models.MenuItem, error) {
	if db.DB == nil {
		return nil, sql.ErrConnDone
	}

	query := "SELECT * FROM menu_items"
	var menuItems []models.MenuItem
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	defer rows.Close()

	rows.Scan(&menuItems)
	for rows.Next() {
		var item models.MenuItem
		if err := rows.Scan(&item.ID, &item.Name, &item.Price); err != nil {
			return nil, err
		}
		menuItems = append(menuItems, item)
	}

	return menuItems, nil
}