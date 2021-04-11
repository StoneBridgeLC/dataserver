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
	e.GET("/news", GetNewsAll)
	e.GET("/news/range", GetNewsWithRange)
	e.GET("/news/:id",GetNewsWithId)

	e.GET("/topic", GetTopicAll)
	e.GET("/topic/range", GetTopicWithRange)
	e.GET("/topic/:id",GetTopicWithId)

	e.GET("/comment", GetCommentAll)
	e.GET("/comment/range", GetCommentWithRange)
	e.GET("/comment/:id",GetCommentWithId)

	e.Logger.Fatal(e.Start(":8080"))
}
