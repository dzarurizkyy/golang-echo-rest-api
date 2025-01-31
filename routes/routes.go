package routes

import (
	"golang-echo-rest-api/controllers"
	"golang-echo-rest-api/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	g := e.Group("")
	g.Use(middleware.IsAuthenticated)

	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "Configuration success!") })
	e.POST("/login", controllers.CheckLogin)
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.GET("/test-struct-validation", controllers.TestStructValidation)

	g.GET("/employee", controllers.GetAllEmployee)
	g.POST("/employee", controllers.AddEmployee)
	g.PUT("/employee/:id", controllers.UpdateEmployee)
	g.DELETE("/employee/:id", controllers.DeleteEmployee)

	return e
}
