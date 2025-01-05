package service

import (
	"errors"
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

func (s *URLService) ShortenURL(originalURL string, shortURL string) error {
	existingURL, _ := s.repo.GetOriginalURL(shortURL)
	if existingURL != "" {
		return errors.New("short URL already exists")
	}

	url := domain.URL{OriginalURL: originalURL, ShortURL: shortURL}
	return s.repo.CreateURL(url)
}

func (s *URLService) ResolveURL(shortURL string) (string, error) {
	return s.repo.GetOriginalURL(shortURL)
}

func (s *URLService) GenerateShortCode() string {
	return generateShortCode()
}

func generateShortCode() string {
	rand.Seed(time.Now().UnixNano()) 

	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	shortCode := make([]rune, 6)

	for i := range shortCode {
		shortCode[i] = chars[rand.Intn(len(chars))]
	}

	return string(shortCode)
}

func (s *URLService) GetAllURLs() ([]domain.URL, error) {
	return s.repo.GetAllURLs()
}
