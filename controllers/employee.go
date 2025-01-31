package controllers

import (
	"golang-echo-rest-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)


func GetAllEmployee(c echo.Context) error {
	result, err := models.GetAllEmployee()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func AddEmployee(c echo.Context) error {
	name := c.FormValue("name")
	address := c.FormValue("address")
	phone_number := c.FormValue("phone_number")

	result, err := models.AddEmployee(name, address, phone_number)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}