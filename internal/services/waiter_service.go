package services

import (
	"database/sql"
	"log"

	"github.com/noBthd/restaurant_api.git/internal/db"
	"github.com/noBthd/restaurant_api.git/internal/models"
)

func CreateWaiter(waiter models.Waiter) error {
	if db.DB == nil {
		log.Print("Database connection is not initialized")
		return sql.ErrConnDone
	}

	query := "INSERT INTO waiters (name, surname) VALUES ($1, $2)"
	_, err := db.DB.Exec(query, waiter.Name, waiter.Surname)
	if err != nil {
		log.Printf("Failed to create waiter: %v", err)
		return err
	}

	return nil
}