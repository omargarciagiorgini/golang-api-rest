package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest/db"
	"go-rest/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

//stgsgrgr
func GetProds(c echo.Context) error {
	catName := c.QueryParam("name")
	return c.String(http.StatusOK, fmt.Sprintf("el nombre del gato es : %s", catName))
}

//InsertProd get json data from body and save it to DB
func InsertProd(c echo.Context) error {
	prod := models.Product{}
	err := json.NewDecoder(c.Request().Body).Decode(&prod)
	if err != nil {
		return err
	}
	pgconn, err := db.Conn()
	pgconn.Create(&prod)
	return c.String(http.StatusOK, "all went fine")
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "hola mundo !!!")
}
