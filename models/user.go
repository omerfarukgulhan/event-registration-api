package models

import (
	"errors"
	"example.com/event-registration-app/db"
	"example.com/event-registration-app/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := "INSERT INTO users(email, password) VALUES ($1, $2)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	user.ID = userId

	return err
}

func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = $1"
	row := db.DB.QueryRow(query, user.Email)
	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)
	if err != nil {
		return errors.New("Credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
