package services

import (
	"database/sql"
	"log"

	"github.com/noBthd/restaurant_api.git/internal/db"
	"github.com/noBthd/restaurant_api.git/internal/models"
)

func CreateMenuOrder(MenuOrder *models.MenuOrder, reservationID int) error {
	if db.DB == nil {
		log.Println("Database connection is nil")
		return sql.ErrConnDone
	}

	// Get the order ID associated with the reservation ID
	query := `SELECT orders.id
		FROM orders
		JOIN bill ON orders.bill_id = bill.id
		JOIN reservations ON bill.reservation_id = reservations.id
		WHERE reservations.id = $1;`
	
	row := db.DB.QueryRow(query, reservationID)
	var orderID int
	
	err := row.Scan(&orderID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No order found for reservation ID %d: %v\n", reservationID, err)
			return err
		}
		log.Printf("Error querying order ID for reservation ID %d: %v\n", reservationID, err)
		return err
	}
	MenuOrder.OrderID = orderID

	// Insert the menu order into the database
	query = "INSERT INTO menu_orders (order_id, menu_item_id, quantitiy) VALUES ($1, $2, $3)"
	_, err = db.DB.Exec(query, MenuOrder.OrderID, MenuOrder.MenuItemID, MenuOrder.Quantity)
	if err != nil {
		log.Printf("Error inserting menu order: %v\n", err)
		return err
	}

	return nil
}