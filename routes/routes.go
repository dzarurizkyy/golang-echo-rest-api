package routes

import (
	"golang-echo-rest-api/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "Configuration success!") })
	e.GET("/employee", controllers.GetAllEmployee)
	e.POST("/employee", controllers.AddEmployee)
	e.PUT("/employee/:id", controllers.UpdateEmployee)
	e.DELETE("/employee/:id", controllers.DeleteEmployee)
	
	return e
}
