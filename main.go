package main

import (
	"fmt"
	"go-rest/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hola mundo !!!")
}

//authentication olahjdoah*/
func authentication(username, password string, c echo.Context) (bool, error) {
	//check against DB
	fmt.Println(username)
	if password == "pass" {
		return true, nil
	}
	return false, nil
}
func main() {

	e := echo.New()
	e.GET("/hello", hello)
	e.Use(middleware.Logger())
	e.Use(middleware.BasicAuth(authentication))
	e.GET("/cats", controllers.GetCats)
	e.Start(":8080")
}
