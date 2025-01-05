package service

import (
	"fiber-starter/internal/domain"
	"fiber-starter/internal/repository"
	"math/rand"
	"time"
)

type URLService struct {
    repo *repository.URLRepository
}

func NewURLService(repo *repository.URLRepository) *URLService {
    return &URLService{repo: repo}
}

func (s *URLService) ShortenURL(originalURL string) (string, error) {
    shortURL := generateShortCode()
    url := domain.URL{OriginalURL: originalURL, ShortURL: shortURL}
    return shortURL, s.repo.CreateURL(url)
}

func (s *URLService) ResolveURL(shortURL string) (string, error) {
    return s.repo.GetOriginalURL(shortURL)
}

func generateShortCode() string {
    rand.Seed(time.Now().UnixNano())
    letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
    code := make([]rune, 6)
    for i := range code {
        code[i] = letters[rand.Intn(len(letters))]
    }
    return string(code)
}
