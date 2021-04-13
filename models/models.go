package models

import (
	"github.com/jmoiron/sqlx"
	"time"
)

const DefaultNewsQueryString = "select * from news "
const DefaultCommentQueryString = "select * from comment "
const DefaultTopicQueryString = "select * from topic "

type News struct {
	Id	int		`json:"id"`
	Body	string	`json:"body"`
	Hash	string	`json:"hash"`
	CreateTime	time.Time	`json:"create_time" db:"create_time"`
	UpdateTime	time.Time	`json:"updated_time" db:"update_time"`
}

type Topic struct {
	Id	int	`json:"id"`
	Topic	string	`json:"topic"`
	Positive	int	`json:"positive"`
	Negative	int	`json:"negative"`
}

type Comment struct {
	Id	int	`json:"id"`
	Nid	int	`json:"nid"`  // New id
	Body	string	`json:"body"`
	Pid		int	`json:"pid"` // Parent id
	IsPos	int	`json:"isPos"`	// This comments sentiment
	CreateTime	time.Time	`json:"create_time"`
	UpdateTime	time.Time	`json:"update_time"`
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

// For select all records
func WithAll() Option {
	return optionFunc(func(o *options) {
	})
}

// For select records in range by id.
func WithRange(from int, to int) Option {
	return optionFunc(func(o *options) {
		o.queryString += "where id >= ? and id <= ?"
		o.args = append(o.args, from, to)
	})
}

// For select record by id
func WithId(id int) Option {
	return optionFunc(func(o *options) {
		o.queryString += "where id = ?"
		o.args = append(o.args, id)
	})
}

// For select topic of news
func TopicWithNews(id int) Option {
	return optionFunc(func(o *options) {
		o.queryString += " as t join on news_topic as nt where nt.nid = ?"
		o.args = append(o.args, id)
	})
}

// For select comments of news.
func CommentWithNews(id int) Option {
	return optionFunc(func(o *options) {
		o.queryString += "where nid = ?"
		o.args = append(o.args, id)
	})
}

// Get news method.
func GetNews(db *sqlx.DB, opts ... Option) ([]News, error) {
	var news []News

	options := options {
		queryString: DefaultNewsQueryString,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	rows, err := db.Queryx(options.queryString, options.args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var currentRow News
		err := rows.StructScan(&currentRow)
		if err != nil {
			return nil, err
		}
		news = append(news, currentRow)
	}

	return news, nil
}

// Get topic method.
func GetTopic(db *sqlx.DB, opts ... Option) ([]Topic, error) {
	var topics []Topic

	options := options {
		queryString: DefaultTopicQueryString,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	rows, err := db.Queryx(options.queryString, options.args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var currentRow Topic
		err := rows.StructScan(&currentRow)
		if err != nil {
			return nil, err
		}
		topics = append(topics, currentRow)
	}

	return topics, nil
}

// Get Comments method.
func GetComment(db *sqlx.DB, opts ... Option) ([]Comment, error) {
	var comments []Comment

	options := options {
		queryString: DefaultCommentQueryString,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	rows, err := db.Queryx(options.queryString, options.args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var currentRow Comment
		err := rows.StructScan(&currentRow)
		if err != nil {
			return nil, err
		}
		comments = append(comments, currentRow)
	}

	return comments, nil
}