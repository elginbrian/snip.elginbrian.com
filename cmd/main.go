package main

import (
	"database/sql"
	"fiber-starter/config"
	"fiber-starter/internal/di"
	"fiber-starter/internal/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
    cfg := config.LoadConfig()

    db, err := sql.Open("postgres", cfg.DatabaseURL)
    if err != nil {
        panic(err)
    }

    container := di.NewContainer(db)
    app := fiber.New()

    routes.SetupRoutes(app, container.URLHandler)

    app.Listen(cfg.Port)
}
