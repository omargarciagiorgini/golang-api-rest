package router

import (
	"go-rest/controllers"
	"go-rest/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//New start a router
func New() *echo.Echo {
	e := echo.New()
	e.Validator = &models.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Logger())
	//e.Use(middleware.BasicAuth(middlewars.Authentication))
	e.GET("/prods", controllers.GetAll)
	e.GET("/prods/:id", controllers.GetProdById)
	e.PUT("/prods", controllers.UpdateProd)
	e.POST("/prods", controllers.InsertProd)
	return e
}
