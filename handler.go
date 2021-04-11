package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const DefaultNewsQueryString = "select * from news"
const DefaultCommentQueryString = "select * from comments"
const DefaultTopicQueryString = "select * from topics"

type News struct {
	Id	int		`json: id`
	Body	string	`json: body`
}

type Topic struct {
	Id	int	`json: id`
	Topic	string	`json: topic`
	Positive	int	`json: positive`
	Negative	int	`json: negative`
}

type Comment struct {
	Id	int	`json: id`
	Nid	int	`json: nid`  // New id
	Body	string	`json: body`
	Pid		int	`json: pid` // Parent id
	IsPos	int	`json: isPos`	// This comments sentiment
}

// For parameter
type options struct {
	queryString string
	args []interface{}
	queryId	int // id of selected record
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithAll() Option {
	return optionFunc(func(o *options) {
	})
}

func WithRange(from int, to int) Option {
	return optionFunc(func(o *options) {
		o.queryString += "where id >= ? and id <= ?"
		o.args = append(o.args, from, to)
	})
}

func WithId(id int) Option {
	return optionFunc(func(o *options) {
		o.queryString += "where id = ?"
		o.args = append(o.args, id)
	})
}

func GetNews(opts ... Option) ([]News, error) {
	var news []News

	options := options {
		queryString: DefaultNewsQueryString,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	rows, err := db.Query(options.queryString, options.args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var currentRow News
		err := rows.Scan(&currentRow.Id, &currentRow.Body)
		if err != nil {
			return nil, err
		}
		news = append(news, currentRow)
	}

	return news, nil
}

// Handler for response all news.
func GetNewsAll (c echo.Context) error {
	// apiserver/news
	news, err := GetNews(WithAll())
	if err != nil {
		return err
	}

	ret, err := json.Marshal(news)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ret)
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

	news, err := GetNews(WithRange(from, to))
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

	news, err := GetNews(WithId(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, news)
}


// Get topic method.
func GetTopic(opts ... Option) ([]Topic, error) {
	var topics []Topic

	options := options {
		queryString: DefaultTopicQueryString,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	rows, err := db.Query(options.queryString, options.args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var currentRow Topic
		err := rows.Scan(&currentRow.Id, &currentRow.Topic, &currentRow.Positive, &currentRow.Negative)
		if err != nil {
			return nil, err
		}
		topics = append(topics, currentRow)
	}

	return topics, nil
}

// Handler for response all topics.
func GetTopicAll (c echo.Context) error {
	// apiserver/comment
	news, err := GetComment(WithAll())
	if err != nil {
		return err
	}

	ret, err := json.Marshal(news)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ret)
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

	news, err := GetComment(WithRange(from, to))
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

	news, err := GetComment(WithId(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, news)
}

// Get Comments method.
func GetComment(opts ... Option) ([]Comment, error) {
	var comments []Comment

	options := options {
		queryString: DefaultCommentQueryString,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	rows, err := db.Query(options.queryString, options.args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var currentRow Comment
		err := rows.Scan(&currentRow.Id, &currentRow.Nid, &currentRow.Body, &currentRow.Pid, &currentRow.IsPos)
		if err != nil {
			return nil, err
		}
		comments = append(comments, currentRow)
	}

	return comments, nil
}

// Handler for response all news.
func GetCommentAll (c echo.Context) error {
	// apiserver/comment
	news, err := GetComment(WithAll())
	if err != nil {
		return err
	}

	ret, err := json.Marshal(news)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ret)
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

	news, err := GetComment(WithRange(from, to))
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

	news, err := GetComment(WithId(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, news)
}