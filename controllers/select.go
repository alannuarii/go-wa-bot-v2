package controllers

import (
	"go-wa-bot-v2/db"
	"go-wa-bot-v2/models"
	"log"
)

// GetAllUser retrieves all users from the database and returns them along with any error encountered.
func GetAllUser() ([]models.User, error) {
	db := db.DB

	users := []models.User{}

	query := `SELECT id_user, name, email, role FROM user`

	err := db.Select(&users, query)
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return nil, err
	}

	return users, nil
}
