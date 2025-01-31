package models

import (
	"database/sql"
	"errors"
	"golang-echo-rest-api/db"
	"golang-echo-rest-api/helpers"
)

type User struct {
	Id    int    `json:"id"`
	Email string `json:"name"`
}

func CheckLogin(email, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE email=$1"

	err := con.QueryRow(sqlStatement, email).Scan(&obj.Id, &obj.Email, &pwd)

	if err == sql.ErrNoRows {
		return false, errors.New("email not found")
	}

	if err != nil {
		return false, errors.New("query Error")
	}

	match, _ := helpers.CheckPasswordHash(password, pwd)

	if !match {
		return false, errors.New("hash and password doesn't match")
	}

	return true, nil
}
