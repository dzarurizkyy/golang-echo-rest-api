package controllers

import (
	"golang-echo-rest-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)


func FetchAllEmployee(c echo.Context) error {
	result, err := models.FetchAllEmployee()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}