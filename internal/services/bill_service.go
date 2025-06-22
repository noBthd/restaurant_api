package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/noBthd/restaurant_api.git/internal/db"
	"github.com/noBthd/restaurant_api.git/internal/models"
)	

func GetAllBills() ([]models.Bill, error) {
	if db.DB == nil {
		log.Println("Database conn is not initialized")
		return nil, sql.ErrConnDone
	}

	query := "SELECT * FROM bill"
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Printf("Error querying bills: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var bills []models.Bill
	for rows.Next() {
		var bill models.Bill
		if err := rows.Scan(&bill.ID, &bill.ReservationID, &bill.TotalPrice, &bill.IsPaid); err != nil {
			log.Printf("Error scanning bill: %v\n", err)
			return nil, err
		}

		bills = append(bills, bill)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over bills: %v\n", err)
		return nil, err
	}

	return bills, nil
}

func GetBillByReservationID(reservationID int) (*models.Bill, error)  {
	if db.DB == nil {
		log.Println("Database conn is not initialized")
		return nil, sql.ErrConnDone
	}

	query := "SELECT * FROM bill WHERE reservation_id = $1"
	row := db.DB.QueryRow(query, reservationID)

	var bill models.Bill
	if err := row.Scan(&bill.ID, &bill.ReservationID, &bill.TotalPrice, &bill.IsPaid); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No bill found for reservation ID %d\n", reservationID)
			return nil, nil
		}
		log.Printf("Error scanning bill: %v\n", err)
		return nil, err
	}

	return &bill, nil
}

func PayBill(reservationID int) error {
	if db.DB == nil {
		log.Println("Database conn is not initialized")
		return sql.ErrConnDone
	}

	query := "SELECT is_paid FROM bill WHERE reservation_id = $1"
	row := db.DB.QueryRow(query, reservationID)
	var isPaid bool
	if err := row.Scan(&isPaid); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No bill found for reservation ID %d\n", reservationID)
			return fmt.Errorf("no bill found for reservation ID %d", reservationID)
		}
		log.Printf("Error scanning bill payment status: %v\n", err)
		return err
	}

	if isPaid {
		log.Printf("Bill for reservation ID %d is already paid\n", reservationID)
		return fmt.Errorf("bill for reservation ID %d is already paid", reservationID)
	}

	query = "UPDATE bill SET is_paid = true WHERE reservation_id = $1"
	_, err := db.DB.Exec(query, reservationID)
	if err != nil {
		log.Printf("Error updating bill payment status: %v\n", err)
		return err
	}

	return nil
}