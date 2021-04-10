package main

import (
	"net/http"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)


func main() {
	// Initialize database
	err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})
}
