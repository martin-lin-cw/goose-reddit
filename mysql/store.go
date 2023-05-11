package mysql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/martin-lin-cw/goose-reddit/goreddit"

	_ "github.com/go-sql-driver/mysql"
)

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	return &Store{
		ThreadStore:  NewThreadStore(db),
		PostStore:    NewPostStore(db),
		CommentStore: NewCommentStore(db),
	}, nil
}

type Store struct {
	goreddit.ThreadStore
	goreddit.PostStore
	goreddit.CommentStore
}
