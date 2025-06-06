package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"todo-api/internal/config"
	"todo-api/internal/handlers"
	"todo-api/internal/repository/db"
)

func main() {

	// Загрузка переменных окружения из .env файла
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// Инициализация базы данных
	if err := config.InitDatabase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	repo := db.NewPostgresRepository(config.DB)
	app := fiber.New()
	handlers.SetupRoutes(app, repo)
	app.Listen(":3000")
}
