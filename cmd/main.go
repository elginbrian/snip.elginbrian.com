package main

import (
	"database/sql"
	"fiber-starter/config"
	"fiber-starter/internal/di"
	"fiber-starter/internal/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	err = config.MigrateDatabase(db)
	if err != nil {
		log.Fatalf("Error applying migrations: %v", err)
	}

	container := di.NewContainer(db)

	app := fiber.New()

	routes.SetupRoutes(app, container.URLHandler)

	err = app.Listen(cfg.Port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	fmt.Println("Server started successfully!")
}
