package services

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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
	
	startTime, err := time.Parse("2006-01-02 15:04:05", reservation.StartTime)
	if err != nil {
		log.Printf("Invalid start time format: %v", err)
		return fmt.Errorf("invalid start time format")
	}
	endTime, err := time.Parse("2006-01-02 15:04:05", reservation.EndTime)
	if err != nil {
		log.Printf("Invalid end time format: %v", err)
		return fmt.Errorf("invalid end time format")
	}
	if startTime.Hour() < 10 || endTime.Hour() > 22 || (endTime.Hour() == 22 && endTime.Minute() > 0) {
		log.Printf("Reservation time is outside of allowed hours (10:00 - 22:00)")
		return fmt.Errorf("reservations are only allowed between 10:00 and 22:00")
	}
	if startTime.Add(30 * time.Minute).After(endTime) {
		log.Printf("Reservation duration must be at least 30 minutes")
		return fmt.Errorf("reservation duration must be at least 30 minutes")
	}
	now := time.Now()
	if !startTime.After(now) || !endTime.After(now) {
		log.Printf("Reservation times must be in the future")
		return fmt.Errorf("reservation times must be in the future")
	}
	

	query := `
		SELECT COUNT(1)
		FROM reservations
		WHERE table_id = $1
			AND is_active = true
			AND (
				($2 < end_time + INTERVAL '30 minutes') AND
				($3 > start_time - INTERVAL '30 minutes') AND
				($2 < end_time AND $3 > start_time)
			)
	`
	var count int
	err = db.DB.QueryRow(query, reservation.TableID, reservation.StartTime, reservation.EndTime).Scan(&count)
	if err != nil {
		log.Printf("Failed to check for overlapping reservations: %v", err)
		return err
	}
	if count > 0 {
		return fmt.Errorf("table %d is already reserved in the selected time range or within 30 minutes before/after", reservation.TableID)
	}

	// Insert the reservation
	insertQuery := `
		INSERT INTO reservations (table_id, user_id, start_time, end_time, is_active)
		VALUES ($1, $2, $3, $4, true)
	`
	_, err = db.DB.Exec(insertQuery, reservation.TableID, reservation.UserID, reservation.StartTime, reservation.EndTime)
	if err != nil {
		log.Printf("Failed to create reservation: %v", err)
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

func GetReservedTableByUserID(userID int) ([]models.Reservation, error) {
	if db.DB == nil {
		log.Print("Database connection is not initialized")
		return nil, sql.ErrConnDone
	}

	query := "SELECT * FROM reservations WHERE user_id = $1"
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		log.Printf("Failed to get reservations for user ID %d: %v", userID, err)
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
			log.Printf("Failed to scan reservation: %v", err)
			return nil, err
		}
		// Check if reservation is active
		if !reservation.Is_active {
			log.Printf("Reservation with ID %d is not active", reservation.ID)
			continue
		}
		reservations = append(reservations, reservation)
	}

	if len(reservations) == 0 {
		log.Printf("No active reservations found for user ID %d", userID)
		return nil, fmt.Errorf("no active reservations found for user ID %d", userID)
	}

	return reservations, nil
}