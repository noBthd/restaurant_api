package services

import (
	"database/sql"
	"log"

	"github.com/noBthd/restaurant_api.git/internal/db"
	"github.com/noBthd/restaurant_api.git/internal/models"
)

func GetAllTables() ([]models.Table, error) {
	if db.DB == nil {
		log.Println("Database connection is not initialized")
		return nil, sql.ErrConnDone
	}

	query := "SELECT * FROM tables"
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Printf("Error querying tables: %v", err)
		return nil, err
	}
	defer rows.Close()

	var tables []models.Table
	for rows.Next() {
		var table models.Table
		if err := rows.Scan(&table.ID, &table.Seats); err != nil {
			log.Printf("Error scanning table: %v", err)
			return nil, err
		}
		tables = append(tables, table)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over tables: %v", err)
		return nil, err
	}

	return tables, nil
}