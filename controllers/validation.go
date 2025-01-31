package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Customer struct {
	Name    string `validate:"required"`
	Email   string `validate:"required,email"`
	Address string `validate:"required"`
	Age     int    `validate:"gte=17,lte=35"`
}

func TestStructValidation(c echo.Context) error {
	v := validator.New()

	cust := Customer{
		Name:    "Dzaru Rizky Fathan Fortuna",
		Email:   "dzarurizkybusiness@gmail.com",
		Address: "Surabaya, Indonesia",
		Age:     22,
	}

	err := v.Struct(cust)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}
