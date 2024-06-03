package models

import "github.com/dexter0323/go-learn/api/db"

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query, err := db.DB.Exec(`
		INSERT INTO users(email, password) 
		VALUES (?, ?)
	`, u.Email, u.Password)
	if err != nil {
		return err
	}

	userId, err := query.LastInsertId()
	u.ID = userId

	return err
}
