package main

import (
	"github.com/StoneBridgeLC/dataserver/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Handler for response all news.
func GetNewsAll (c echo.Context) error {
	// apiserver/news
	news, err := models.GetNews(db, models.WithAll())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, news)
}

// Handler for response news in range.
func GetNewsWithRange (c echo.Context) error {
	// apiserver/news/ragne?from=11?to=11
	from, err := strconv.Atoi(c.QueryParam("from"))
	if err != nil {
		return err
	}
	to, err := strconv.Atoi(c.QueryParam("to"))
	if err != nil {
		return err
	}

	news, err := models.GetNews(db, models.WithRange(from, to))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, news)
}

// Handler for response news with id
func GetNewsWithId (c echo.Context) error {
	// apiserver/news?id=123123
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	news, err := models.GetNews(db, models.WithId(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, news)
}

// Handler for response all topics.
func GetTopicAll (c echo.Context) error {
	// apiserver/comment
	topics, err := models.GetComment(db, models.WithAll())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, topics)
}

// Handler for response topics in range.
func GetTopicWithRange (c echo.Context) error {
	// apiserver/comment/range?from=11?to=11
	from, err := strconv.Atoi(c.QueryParam("from"))
	if err != nil {
		return err
	}
	to, err := strconv.Atoi(c.QueryParam("to"))
	if err != nil {
		return err
	}

	news, err := models.GetComment(db, models.WithRange(from, to))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, news)
}

// Handler for response topic with id
func GetTopicWithId (c echo.Context) error {
	// apiserver/comment?id=123123
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	news, err := models.GetComment(db, models.WithId(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, news)
}

// Handler for response all news.
func GetCommentAll (c echo.Context) error {
	// apiserver/comment
	comments, err := models.GetComment(db, models.WithAll())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, comments)
}

// Handler for response news in range.
func GetCommentWithRange (c echo.Context) error {
	// apiserver/comment/range?from=12?to=12
	from, err := strconv.Atoi(c.QueryParam("from"))
	if err != nil {
		return err
	}
	to, err := strconv.Atoi(c.QueryParam("to"))
	if err != nil {
		return err
	}

	news, err := models.GetComment(db, models.WithRange(from, to))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, news)
}

// Handler for response news with id
func GetCommentWithId (c echo.Context) error {
	// apiserver/comment/123123
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	news, err := models.GetComment(db, models.WithId(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, news)
}

// Get Topic of news
func GetTopicOfNews (c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	topics, err := models.GetTopic(db, models.TopicWithNews(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, topics)
}

// Get Comments of news
func GetCommentOfNews (c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	comments, err := models.GetComment(db, models.CommentWithNews(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, comments)
}