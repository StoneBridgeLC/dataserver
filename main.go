package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize database
	err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "remote_ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, time=${time_custom}, error=${error}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})

	e.GET("/news", GetNewsAll)
	e.GET("/news/range", GetNewsWithRange)
	e.GET("/news/:id",GetNewsWithId)
	e.GET("/news/month", GetNewsInMonth)

	e.GET("/topic", GetTopicAll)
	e.GET("/topic/range", GetTopicWithRange)
	e.GET("/topic/:id",GetTopicWithId)
	e.GET("/news/:id/topic", GetTopicOfNews)

	e.GET("/comment", GetCommentAll)
	e.GET("/comment/range", GetCommentWithRange)
	e.GET("/comment/:id",GetCommentWithId)
	e.GET("/news/:id/comment", GetCommentOfNews)

	e.Logger.Fatal(e.Start(":8080"))
}
