package models

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"time"
)

const DefaultNewsQueryString = "select * from news "
const DefaultCommentQueryString = "select * from comment "
const DefaultTopicQueryString = "select * from topic "

type News struct {
	Id	int		`json:"id"`
	Title	string	`json:"title"`
	Body	string	`json:"body"`
	Hash	string	`json:"hash"`
	Url		sql.NullString	`json:"url"`
	CreateTime	time.Time	`json:"create_time" db:"create_time"`
	UpdateTime	time.Time	`json:"update_time" db:"update_time"`
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
	Pid		sql.NullInt64	`json:"pid"` // Parent id
	IsPos	sql.NullInt64	`json:"is_pos" db:"is_pos"`	// This comments sentiment
	CreateTime	time.Time	`json:"create_time" db:"create_time"`
	UpdateTime	time.Time	`json:"update_time" db:"update_time"`
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

func WithInMonth() Option {
	return optionFunc(func(o *options) {
		o.queryString += "where update_time between DATE_ADD(NOW(),INTERVAL -1 MONTH ) AND NOW()"
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

func GetUnlabeledComments(db *sqlx.DB, limit int) ([]Comment, error) {
	rows, err := db.Queryx(`
		select c.id, nid, body, pid, is_pos, create_time, update_time
		from comment c
		where c.is_pos is null
		limit ?;`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := make([]Comment, 0, limit)
	for rows.Next() {
		comment := Comment{}
		if err := rows.StructScan(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func UpdateCommentLabel(db *sqlx.DB, comment Comment) error {
	_, err := db.NamedExec("update comment c set c.is_pos = :is_pos where c.id = :id and c.is_pos is null;", comment)
	return err
}