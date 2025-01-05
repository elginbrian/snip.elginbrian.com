package di

import (
	"database/sql"
	"fiber-starter/internal/handler"
	"fiber-starter/internal/repository"
	"fiber-starter/internal/service"
)

type Container struct {
    URLHandler *handler.URLHandler
}

func NewContainer(db *sql.DB) *Container {
    // Initialize repository
    urlRepo := repository.NewURLRepository(db)

    // Initialize service
    urlService := service.NewURLService(urlRepo)

    // Initialize handler
    urlHandler := handler.NewURLHandler(urlService)

    return &Container{
        URLHandler: urlHandler,
    }
}
