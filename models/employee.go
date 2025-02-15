package models

import (
	"golang-echo-rest-api/db"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Employee struct {
	Id          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

func GetAllEmployee() (Response, error) {
	var obj Employee
	var arrobj []Employee
	var res Response

	con := db.CreateCon()

	rows, err := con.Query("SELECT * FROM employees")
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
	res.Message = "success"
	res.Data = arrobj

	return res, nil
}

func AddEmployee(name, address, phone_number string) (Response, error) {
	var res Response
	var lastInsertedId int

	v := validator.New()

	employee := Employee{
		Name: name,
		Address: address,
		PhoneNumber: phone_number,
	} 

	err := v.Struct(employee)

	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	err = con.QueryRow("INSERT INTO employees (name, address, phone_number) VALUES ($1, $2, $3) RETURNING id", name, address, phone_number).Scan(&lastInsertedId)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]int{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdateEmployee(id int, name, address, phone_number string) (Response, error) {
	var res Response

	con := db.CreateCon()

	result, err := con.Exec("UPDATE employees SET name=$1, address=$2, phone_number=$3 WHERE id=$4", name, address, phone_number, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteEmployee(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	stmt, err := con.Prepare("DELETE FROM employees where id=$1")
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
