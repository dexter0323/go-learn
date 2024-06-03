package models

import (
	"github.com/dexter0323/go-learn/api/db"
	"github.com/dexter0323/go-learn/api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	query, err := db.DB.Exec(`
		INSERT INTO users(email, password) 
		VALUES (?, ?)
	`, u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := query.LastInsertId()
	u.ID = userId

	return err
}
