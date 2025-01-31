package models

import (
	"golang-echo-rest-api/db"
	"net/http"
)

type Employee struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

func GetAllEmployee() (Response, error) {
	var obj Employee
	var arrobj []Employee
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM employee"
	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Name, &obj.Address, &obj.PhoneNumber, &obj.Id)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func AddEmployee(name, address, phone_number string) (Response, error) {
	var res Response
	var lastInsertedId int

	con := db.CreateCon()

	err := con.QueryRow("INSERT INTO employee (name, address, phone_number) VALUES ($1, $2, $3) RETURNING id", name, address, phone_number).Scan(&lastInsertedId)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}
