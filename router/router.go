package router

import (
	"go-rest/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//New start a router
func New() *echo.Echo {
	e := echo.New()
	e.GET("/hello", controllers.Hello)
	e.Use(middleware.Logger())
	//e.Use(middleware.BasicAuth(middlewars.Authentication))
	e.GET("/prods", controllers.GetProds)
	e.POST("/prods", controllers.InsertProd)
	return e
}
