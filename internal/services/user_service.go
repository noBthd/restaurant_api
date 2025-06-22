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

func CreateUser(user *models.User) error {
	if db.DB == nil {
		log.Print("Database connection is not initialized")
		return sql.ErrConnDone
	}

	query := "INSERT INTO users (email, password) VALUES ($1, $2)"
	_, err := db.DB.Exec(query, user.Email, user.Password)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return err
	}

	return nil
}

func LoginUser(user *models.User) (*models.User, error) {
	if db.DB == nil {
		log.Print("Database connection is not initialized")
		return nil, sql.ErrConnDone
	}

	query := "SELECT id, email, password, creation_date, is_admin FROM users WHERE email = $1 AND password = $2"
	row := db.DB.QueryRow(query, user.Email, user.Password)

	var existingUser models.User
	err := row.Scan(&existingUser.ID, &existingUser.Email, &existingUser.Password, &existingUser.Creation_date, &existingUser.Is_admin)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user found with email: %s", user.Email)
			return nil, nil
		}

		log.Printf("Failed to login user: %v", err)
		return nil, err
	}

	log.Printf("User logged in successfully: %s", existingUser.Email)
	return &existingUser, nil
}

func RemoveUser(id int) error {
	if db.DB == nil {
		log.Print("Database connection is not initialized")
		return sql.ErrConnDone
	}

	query := "DELETE FROM users WHERE id = $1"
	_, err := db.DB.Exec(query, id)
	if err != nil {
		log.Printf("Failed to remove user with ID %d: %v", id, err)
		return err
	}

	log.Printf("User with ID %d removed successfully", id)
	return nil
}

func MakeUserAdmin(id int) error {
	if db.DB == nil {
		log.Print("Database connection is not initialized")
		return sql.ErrConnDone
	}

	query := "UPDATE users SET is_admin = TRUE WHERE id = $1"
	_, err := db.DB.Exec(query, id)
	if err != nil {
		log.Printf("Failed to make user with ID %d an admin: %v", id, err)
		return err
	}

	log.Printf("User with ID %d is now an admin", id)
	return nil
}