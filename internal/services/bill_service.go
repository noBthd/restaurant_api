package services

import (
	"database/sql"
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

func PayBill(billID int) error {
	if db.DB == nil {
		log.Println("Database conn is not initialized")
		return sql.ErrConnDone
	}

	query := "UPDATE bill SET is_paid = true WHERE reservation_id = $1"
	_, err := db.DB.Exec(query, billID)
	if err != nil {
		log.Printf("Error updating bill payment status: %v\n", err)
		return err
	}

	return nil
}