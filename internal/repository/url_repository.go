package repository

import (
	"database/sql"
	"fiber-starter/internal/domain"
)

type URLRepository struct {
    db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
    return &URLRepository{db: db}
}

func (r *URLRepository) CreateURL(url domain.URL) error {
    _, err := r.db.Exec("INSERT INTO urls (original_url, short_url) VALUES ($1, $2)", url.OriginalURL, url.ShortURL)
    return err
}

func (r *URLRepository) GetOriginalURL(shortURL string) (string, error) {
    var originalURL string
    err := r.db.QueryRow("SELECT original_url FROM urls WHERE short_url = $1", shortURL).Scan(&originalURL)
    return originalURL, err
}
