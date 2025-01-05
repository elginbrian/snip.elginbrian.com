package handler

import (
	"fiber-starter/internal/service"

	"github.com/gofiber/fiber/v2"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(service *service.URLService) *URLHandler {
	return &URLHandler{service: service}
}

func (h *URLHandler) ShortenURL(c *fiber.Ctx) error {
	var payload struct {
		OriginalURL string `json:"original_url"`
		CustomShortURL string `json:"custom_short_url,omitempty"`
	}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	shortURL := payload.CustomShortURL
	if shortURL == "" {
		shortURL = h.service.GenerateShortCode() 
	}

	err := h.service.ShortenURL(payload.OriginalURL, shortURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not shorten URL"})
	}

	return c.JSON(fiber.Map{"short_url": shortURL})
}

func (h *URLHandler) ResolveURL(c *fiber.Ctx) error {
	shortURL := c.Params("shortURL")
	originalURL, err := h.service.ResolveURL(shortURL)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "URL not found"})
	}

	return c.Redirect(originalURL)
}

func (h *URLHandler) GetAllURLs(c *fiber.Ctx) error {
	urls, err := h.service.GetAllURLs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch URLs"})
	}

	var response []fiber.Map
	for _, url := range urls {
		response = append(response, fiber.Map{
			"original_url": url.OriginalURL,
			"short_url":    "https://snip.elginbrian.com/" + url.ShortURL, 
		})
	}

	return c.JSON(fiber.Map{"urls": response})
}