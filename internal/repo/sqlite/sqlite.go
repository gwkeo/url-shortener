package sqlite

import (
	"database/sql"
	"errors"
	"github.com/gwkeo/url-shortener/internal/repo"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

type Repo struct {
	db *sql.DB
}

func New(dbPath string) (*Repo, error) {

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS urls (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			url TEXT NOT NULL,
			shortened_url TEXT UNIQUE NOT NULL)`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, err
	}

	return &Repo{db: db}, nil
}

func (r *Repo) Create(url, shortened string) error {
	stmt, err := r.db.Prepare(`INSERT INTO urls (url, shortened_url) VALUES (?, ?)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(url, shortened)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) URL(shortURL string) (string, error) {

	stmt, err := r.db.Prepare(`SELECT url FROM urls WHERE shortened_url = ?`)
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
