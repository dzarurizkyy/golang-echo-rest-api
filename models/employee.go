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

func FetchAllEmployee() (Response, error) {
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
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Address, &obj.PhoneNumber)

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
