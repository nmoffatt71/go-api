package models

import (
	"fmt"

	"rest-api.com/m/v2/db"
	"rest-api.com/m/v2/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users (email, password)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err

}
