package routes

import (
	"fiber-starter/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, urlHandler *handler.URLHandler) {
    app.Post("/shorten", urlHandler.ShortenURL)
    app.Get("/:shortURL", urlHandler.ResolveURL)
	app.Get("/", urlHandler.GetAllURLs)
}
