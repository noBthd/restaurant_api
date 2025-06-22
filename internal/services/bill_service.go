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