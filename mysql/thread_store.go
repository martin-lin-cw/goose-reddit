package mysql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/martin-lin-cw/goose-reddit/goreddit"
)

func NewThreadStore(db *sqlx.DB) *ThreadStore {
	return &ThreadStore{
		DB: db,
	}
}

type ThreadStore struct {
	*sqlx.DB
}

func (s *ThreadStore) Thread(id int64) (goreddit.Thread, error) {
	var t goreddit.Thread
	if err := s.Get(&t, `SELECT * FROM thread WHERE id = $1`, id); err != nil {
		return goreddit.Thread{}, fmt.Errorf("error getting thread: %w", err)
	}
	return t, nil
}

func (s *ThreadStore) Threads() ([]goreddit.Thread, error) {
	var tt []goreddit.Thread
	if err := s.Select(&tt, `SELECT * FROM thread`); err != nil {
		return []goreddit.Thread{}, fmt.Errorf("error getting threads: %w", err)
	}
	return tt, nil
}

func (s *ThreadStore) CreateThread(t *goreddit.Thread) error {
	res, err := s.Exec(`INSERT INTO thread (title, description) VALUES ($1, $2)`, t.Title, t.Description)
	if err != nil {
		return fmt.Errorf("error creating thread: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting created id: %w", err)
	}
	t.Id = id
	return nil
}

func (s *ThreadStore) UpdateThread(t *goreddit.Thread) error {
	_, err := s.Exec(`UPDATE thread SET title = $1, description = $2 WHERE id = $3`, t.Title, t.Description, t.Id)
	if err != nil {
		return fmt.Errorf("error updating thread: %w", err)
	}
	return nil
}

func (s *ThreadStore) DeleteThread(id int64) error {
	if _, err := s.Exec(`DELETE FROM thread WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting thread: %w", err)
	}
	return nil
}
