package services

import (
	"database/sql"
	"log"

	"github.com/noBthd/restaurant_api.git/internal/db"
	"github.com/noBthd/restaurant_api.git/internal/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	if db.DB == nil {
		log.Print("Database connection is not initialized")
		return nil, sql.ErrConnDone
	}
	query := "SELECT id, email, password, creation_date, is_admin FROM users WHERE email = $1"
	row := db.DB.QueryRow(query, email)

	var user models.User

	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Creation_date, &user.Is_admin)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user found with email: %s", email)
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}