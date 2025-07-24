package models

import (
	"context"
	"errors"
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
	// QueryRow returns the ROW, on which you can apply Scan method to get the value
	// Then when you apply Scan method on tis row, it may return err, hence you write err
	err := m.DB.QueryRow(context.Background(), query, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets 
			  WHERE expires > CURRENT_TIMESTAMP AND id = $1`
	// var snippet Snippet

	snippet := &Snippet{}

	row := m.DB.QueryRow(context.Background(), query, id)

	err := row.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return snippet, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
