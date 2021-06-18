package main

import (
	"database/sql"
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

// Get News that updated in 30days.
func GetNewsInMonth (c echo.Context) error {
	news, err := models.GetNews(db, models.WithInMonth())
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

// Handler for response all comments.
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

const (
	maxLimit = 1000
	defaultLimit = 100
)

func GetCommentUnlabeled (c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 0 || limit > maxLimit {
		limit = defaultLimit
	}
	if limit < 0 {
		limit = 0
	} else if limit > maxLimit {
		limit = maxLimit
	}

	comments, err := models.GetUnlabeledComments(db, limit)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, comments)
}

func UpdateCommentLabel (c echo.Context) error {
	type body struct {
		IsPos int `json:"is_pos"`
		models.Comment
	}

	var comments []body
	if err := c.Bind(&comments); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	go func(comments []body) {
		for _, comment := range comments {
			comment.Comment.IsPos = sql.NullInt64{
				Int64: int64(comment.IsPos),
				Valid: true,
			}
			if err := models.UpdateCommentLabel(db, comment.Comment); err != nil {
				c.Logger().Errorf("Failed to update comment (comment id %d) : %v", comment.Id, err)
			}
		}
	}(comments)

	return nil
}