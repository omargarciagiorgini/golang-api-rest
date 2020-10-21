package db

import (
	"go-rest/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Conn is the conexion to DB
func Conn() (*gorm.DB, error) {

	dsn := "user=omar password=omar4072 dbname=golang port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	return conn, err
}

//RunMigrations
func RunMigrations(conn *gorm.DB) error {
	conn.AutoMigrate(&models.Product{})
	return nil
}
