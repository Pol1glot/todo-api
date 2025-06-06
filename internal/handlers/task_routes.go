package handlers

import (
	"github.com/gofiber/fiber/v2"
	"todo-api/internal/models"
	"todo-api/internal/repository"
)

func SetupRoutes(app *fiber.App, repo repository.TaskRepository) {
	api := app.Group("/tasks")
	api.Post("/", createTask(repo))
	api.Get("/", getTasks(repo))
	api.Put("/:id", updateTask(repo))
	api.Delete("/:id", deleteTask(repo))
}

func createTask(repo repository.TaskRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		task := new(models.Task)
		if err := c.BodyParser(task); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}

		if err := repo.CreateTask(c.Context(), task); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create task"})
		}

		return c.Status(fiber.StatusCreated).JSON(task)
	}
}

func getTasks(repo repository.TaskRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tasks, err := repo.GetTasks(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve tasks"})
		}

		return c.JSON(tasks)
	}
}

func updateTask(repo repository.TaskRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		task := new(models.Task)
		if err := c.BodyParser(task); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}

		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
		}
		task.ID = id

		if err := repo.UpdateTask(c.Context(), task); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update task"})
		}

		return c.JSON(task)
	}
}

func deleteTask(repo repository.TaskRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
		}

		if err := repo.DeleteTask(c.Context(), id); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete task"})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
