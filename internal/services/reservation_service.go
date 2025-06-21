package services

import (
	"database/sql"
	"log"

	"github.com/noBthd/restaurant_api.git/internal/db"
	"github.com/noBthd/restaurant_api.git/internal/models"
)

func GetAllReservations() ([]models.Reservation, error) {
	if db.DB == nil {
		return nil, sql.ErrConnDone
	}

	query := "SELECT * FROM reservations"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []models.Reservation

	for rows.Next() {
		var reservation models.Reservation
		err := rows.Scan(
			&reservation.ID, 
			&reservation.TableID, 
			&reservation.UserID, 
			&reservation.StartTime, 
			&reservation.EndTime, 
			&reservation.Is_active,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("No reservations found")
				return nil, nil 
			}
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	return reservations, nil
}