package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

//stgsgrgr
func GetCats(c echo.Context) error {
	catName := c.QueryParam("name")
	return c.String(http.StatusOK, fmt.Sprintf("el nombre del gato es : %s", catName))
}
