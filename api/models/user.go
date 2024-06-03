package models

import (
	"errors"
	"fmt"

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

func (u *User) ValidateCredentials() error {
	query := db.DB.QueryRow(`
		SELECT id, password FROM users WHERE email = ?
	`, u.Email)

	var retrievedPassword string
	err := query.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		fmt.Println(err)
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	fmt.Println(passwordIsValid, u.Password, retrievedPassword, err)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}
