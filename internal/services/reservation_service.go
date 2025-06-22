package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/noBthd/restaurant_api.git/internal/db"
	"github.com/noBthd/restaurant_api.git/internal/models"
)

func GetAllReservations() ([]models.Reservation, error) {
	if db.DB == nil {
		log.Print("Database connection is not initialized")
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

func CreateReservation(reservation models.Reservation) error {
	if db.DB == nil {
		log.Print("Database connection is not initialized")
		return sql.ErrConnDone
	}

	query := "INSERT INTO reservations (table_id, user_id, start_time, end_time) VALUES ($1, $2, $3, $4)"
	_, err := db.DB.Exec(query, reservation.TableID, reservation.UserID, reservation.StartTime, reservation.EndTime)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return err
	}

	return nil
}

func CancelReservation(reservationID int) error {
	if db.DB == nil {
		log.Print("Database connection is not initialized")
		return sql.ErrConnDone
	}

	query := "SELECT is_active FROM reservations WHERE id = $1"
	row := db.DB.QueryRow(query, reservationID)
	if row.Err() != nil {
		log.Printf("Failed to find reservation with ID %d: %v", reservationID, row.Err())
		return row.Err()
	}

	var isActive bool
	err := row.Scan(&isActive)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No reservation found with ID %d", reservationID)
			return fmt.Errorf("no reservation found with ID %d", reservationID)
		}
		log.Printf("Failed to scan reservation: %v", err)
		return err
	}
	
	if !isActive {
		log.Printf("Reservation with ID %d is already inactive", reservationID)
		return fmt.Errorf("reservation with ID %d is already inactive", reservationID)
	}

	query = "UPDATE reservations SET is_active = false WHERE id = $1"
	_, err = db.DB.Exec(query, reservationID)
	if err != nil {
		log.Printf("Failed to cancel reservation: %v", err)
		return err
	}

	return nil
}
