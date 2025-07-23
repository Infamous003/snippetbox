package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *pgx.Conn
}

func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	query := `INSERT INTO snippets (title, content, created, expires)
			  VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_DATE + $3 * INTERVAL '1 day')
			  RETURNING id`
	var id int
	err := m.DB.QueryRow(context.Background(), query, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
