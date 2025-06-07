package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"todo-api/internal/config"
	"todo-api/internal/handlers"
	"todo-api/internal/repository/db"
	"todo-api/internal/services"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	requiredEnvVars := []string{"DB_USER", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_NAME", "APP_PORT"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("Required environment variable %s is not set", envVar)
		}
	}

	dbPort := os.Getenv("DB_PORT")
	if _, err := fmt.Sscanf(dbPort, "%d", new(int)); err != nil {
		log.Fatalf("Invalid DB_PORT value: %s", dbPort)
	}

	appPort := os.Getenv("APP_PORT")
	if _, err := fmt.Sscanf(appPort, "%d", new(int)); err != nil {
		log.Fatalf("Invalid APP_PORT value: %s", appPort)
	}

	if err := config.InitDatabase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	repo := db.NewPostgresRepository(config.DB)
	taskService := services.NewTaskService(repo)
	app := fiber.New()
	handlers.SetupRoutes(app, taskService)

	log.Printf("Server is starting on port %s", appPort)
	if err := app.Listen(":" + appPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
