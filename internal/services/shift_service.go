package services

import (
	"database/sql"
	"log"

	"github.com/noBthd/restaurant_api.git/internal/db"
	"github.com/noBthd/restaurant_api.git/internal/models"
)

func CreateShift(shift *models.Shift, tableID int) error {
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

	query = `INSERT INTO table_shifts (table_id, shift_id) VALUES ($1, $2)`
	_, err = db.DB.Exec(query, tableID, shift.ID)
	if err != nil {
		log.Printf("Error creating shift: %v", err)
		return nil
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

func AddServedTableByReservation(reservationID int) error {
	if db.DB == nil {
		log.Println("Database connection is not initialized")
		return sql.ErrConnDone
	}

	query := `UPDATE shifts
				SET tables_served = tables_served + 1
				WHERE id IN (
					SELECT ts.shift_id
					FROM reservations r
					JOIN table_shifts ts ON r.table_id = ts.table_id
					JOIN shifts s ON ts.shift_id = s.id
					WHERE r.id = $1 AND s.date = CURRENT_DATE
				);`
	_, err := db.DB.Exec(query, reservationID)
	if err != nil {
		log.Printf("Error adding served table by reservation: %v", err)
		return err
	}
	return nil
}