package main

import (
	"go-rest/db"
	"go-rest/router"
)

func main() {
	e := router.New()
	pgconn, err := db.Conn()
	if err != nil {
		panic(err)
	}
	db.RunMigrations(pgconn)
	e.Start(":8080")
}
