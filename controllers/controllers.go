package controllers

import (
	"encoding/json"
	"go-rest/db"
	"go-rest/models"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

//stgsgrgr
func GetAll(c echo.Context) error {
	pgconn, err := db.Conn()
	if err != nil {
		return err
	}
	prod := models.Product{}
	pgconn.Find(&prod)
	return c.JSON(http.StatusOK, prod)
}

//InsertProd get json data from body and save it to DB
func InsertProd(c echo.Context) (err error) {
	pgconn, err := db.Conn()
	prod := new(models.Product)
	if err = c.Bind(prod); err != nil {
		return
	}
	if err = c.Validate(prod); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, strings.ToLower(err.Error()))
	}
	pgconn.Create(&prod)
	return c.JSON(http.StatusOK, prod)
}

//fafaf fas
func GetProdById(c echo.Context) error {
	pgconn, err := db.Conn()
	if err != nil {
		return err
	}
	prod := models.Product{}
	id := c.Param("id")
	pgconn.Find(&prod, id)
	return c.JSON(http.StatusOK, prod)
}

//opjpohoh
func UpdateProd(c echo.Context) error {
	pgconn, err := db.Conn()
	prod := models.Product{}
	if err != nil {
		return err
	}
	err = json.NewDecoder(c.Request().Body).Decode(&prod)
	if err != nil {
		return err
	}
	pgconn.Save(&prod)

	return c.JSON(http.StatusOK, prod)
}
