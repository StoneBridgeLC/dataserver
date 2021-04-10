package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const DefaultNewsQueryString = "select * from news"

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

func GetNewsWithRange (c echo.Context) error {
	// apiserver/news?from=11?to=11
	from, err := strconv.Atoi(c.Param("from"))
	if err != nil {
		return err
	}
	to, err := strconv.Atoi(c.Param("to"))
	if err != nil {
		return err
	}

	news, err := GetNews(WithRange(from, to))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, news)
}

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