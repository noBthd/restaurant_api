package services

import (
	"database/sql"
	"log"

	"github.com/noBthd/restaurant_api.git/internal/db"
	"github.com/noBthd/restaurant_api.git/internal/models"
)

func CreateShift(shift *models.Shift) error {
	if db.DB == nil {
		log.Println("Database connection is not initialized")
		return sql.ErrConnDone
	}

	query := `INSERT INTO shifts (waiter_id, date, tables_served) VALUES ($1, $2, $3) RETURNING id`
	err := db.DB.QueryRow(query, shift.WaiterID, shift.Date, shift.TablesServed).Scan(&shift.ID)
	if err != nil {
		log.Printf("Error creating shift: %v", err)
		return err
	}
	return nil
}

func AddServedTable(shiftID int) error {
	if db.DB == nil {
		log.Println("Database connection is not initialized")
		return sql.ErrConnDone
	}

	query := "UPDATE shifts SET tables_served = tables_served + 1 WHERE id = $1"
	_, err := db.DB.Exec(query, shiftID)
	if err != nil {
		log.Printf("Error adding served table: %v", err)
		return err
	}
	return nil
}