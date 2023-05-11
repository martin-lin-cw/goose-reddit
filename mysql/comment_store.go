package mysql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/martin-lin-cw/goose-reddit/goreddit"
)

func NewCommentStore(db *sqlx.DB) *CommentStore {
	return &CommentStore{DB: db}
}

type CommentStore struct {
	*sqlx.DB
}

func (s *CommentStore) Comment(id int64) (goreddit.Comment, error) {
	var c goreddit.Comment
	if err := s.Get(&c, `SELECT * FROM comment WHERE id = $1`, id); err != nil {
		return goreddit.Comment{}, fmt.Errorf("error getting comment: %w", err)
	}
	return c, nil
}

func (s *CommentStore) CommentsByPost(postId int64) ([]goreddit.Comment, error) {
	var cc []goreddit.Comment
	if err := s.Select(&cc, `SELECT * FROM comment WHERE post_id = $1`, postId); err != nil {
		return []goreddit.Comment{}, fmt.Errorf("error getting comments by postId: %w", err)
	}
	return cc, nil
}

func (s *CommentStore) CreateComment(c *goreddit.Comment) error {
	res, err := s.Exec(`INSERT INTO comment (post_id, content, votes) VALUES ($1, $2, $3)`,
		c.PostId, c.Content, c.Votes)
	if err != nil {
		return fmt.Errorf("error creating comment: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting created comment id:%w", err)
	}
	c.Id = id
	return nil
}

func (s *CommentStore) UpdateComment(c *goreddit.Comment) error {
	_, err := s.Exec(`UPDATE comment SET post_id = $1, content = $2, votes = $3 WHERE id = $4`,
		c.PostId, c.Content, c.Votes, c.Id)
	if err != nil {
		return fmt.Errorf("error updating comment: %w", err)
	}
	return nil
}

func (s *CommentStore) DeleteComment(id int64) error {
	_, err := s.Exec(`DELETE FROM comment WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting comment: %w", err)
	}
	return nil
}
