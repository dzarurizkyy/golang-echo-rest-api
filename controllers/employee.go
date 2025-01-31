package controllers

import (
	"golang-echo-rest-api/models"
	"net/http"
	"strconv"

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
	phoneNumber := c.FormValue("phone_number")

	result, err := models.AddEmployee(name, address, phoneNumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateEmployee(c echo.Context) error {
	id := c.Param("id")
	name := c.FormValue("name")
	address := c.FormValue("address")
	phoneNumber := c.FormValue("phone_number")

	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"messsage": err.Error()})
	}

	result, err := models.UpdateEmployee(convId, name, address, phoneNumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"messsage": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}