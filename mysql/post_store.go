package mysql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/martin-lin-cw/goose-reddit/goreddit"
)

func NewPostStore(db *sqlx.DB) *PostStore {
	return &PostStore{DB: db}
}

type PostStore struct {
	*sqlx.DB
}

func (s *PostStore) Post(id int64) (goreddit.Post, error) {
	var p goreddit.Post
	if err := s.Get(&p, `SELECT * FROM post WHERE id = $1`, id); err != nil {
		return goreddit.Post{}, fmt.Errorf("error getting post %w", err)
	}
	return p, nil
}

func (s *PostStore) PostsByThread(threadId int64) ([]goreddit.Post, error) {
	var pp []goreddit.Post
	if err := s.Select(&pp, `SELECT * FROM post WHERE thread_id = $1`, threadId); err != nil {
		return []goreddit.Post{}, fmt.Errorf("error getting posts by thread: %w", err)
	}
	return pp, nil
}

func (s *PostStore) CreatePost(p *goreddit.Post) error {
	res, err := s.Exec(`INSERT INTO post (thread_id, title, content, votes) VALUES ($1, $2, $3, $4)`,
		p.ThreadId, p.Title, p.Content, p.Votes)
	if err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting created id: %w", err)
	}
	p.Id = id
	return nil
}

func (s *PostStore) UpdatePost(p *goreddit.Post) error {
	_, err := s.Exec(`UPDATE post SET thread_id = $1, title = $2 content = $3, votes = $4 WHERE id = $5`,
		p.ThreadId, p.Title, p.Content, p.Votes, p.Id)
	if err != nil {
		return fmt.Errorf("error updating post: %w", err)
	}
	return nil
}

func (s *PostStore) DeletePost(id int64) error {
	_, err := s.Exec(`DELETE FROM post WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting post: %w", err)
	}
	return nil
}
