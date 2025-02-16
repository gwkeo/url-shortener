package sqlite

import (
	"database/sql"
	"errors"
	"github.com/gwkeo/url-shortener/internal/repo"
)

type Repo struct {
	db *sql.DB
}

func New() *Repo { return &Repo{db: nil} }

func (r *Repo) Create(url, shortened string) (int64, error) {
	stmt, err := r.db.Prepare(`INSERT INTO urls (origin, shortened) VALUES (?, ?)`)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(url, shortened)
	if err != nil {
		return 0, err
	}

	shortenedUrl, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return shortenedUrl, nil
}

func (r *Repo) URL(shortURL string) (string, error) {

	stmt, err := r.db.Prepare(`SELECT origin FROM urls WHERE shortened = ?`)
	if err != nil {
		return "", err
	}

	row := stmt.QueryRow(shortURL)
	var origin string
	if err = row.Scan(&origin); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", repo.ErrInvalidShortURL
		}
		return "", err
	}

	return origin, nil
}
