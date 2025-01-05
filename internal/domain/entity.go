package domain

type URL struct {
	ID          int    `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	CreatedAt   string `json:"created_at"`
}
