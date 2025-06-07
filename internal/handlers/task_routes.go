package handlers

import (
	"github.com/gofiber/fiber/v2"
	"todo-api/internal/api"
	"todo-api/internal/dtos"
)

type TaskHandler struct {
	service api.TaskService
}

func NewTaskHandler(service api.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func SetupRoutes(app *fiber.App, service api.TaskService) {
	handler := NewTaskHandler(service)
	api := app.Group("/tasks")
	api.Post("/", handler.CreateTask)
	api.Get("/", handler.GetTasks)
	api.Put("/:id", handler.UpdateTask)
	api.Delete("/:id", handler.DeleteTask)
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	var createDTO dtos.CreateTaskDTO
	if err := c.BodyParser(&createDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	task, err := h.service.CreateTask(c.Context(), &createDTO)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create task"})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {
	tasks, err := h.service.GetTasks(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve tasks"})
	}

	return c.JSON(tasks)
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	var updateDTO dtos.UpdateTaskDTO
	if err := c.BodyParser(&updateDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	task, err := h.service.UpdateTask(c.Context(), id, &updateDTO)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update task"})
	}

	return c.JSON(task)
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	err = h.service.DeleteTask(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete task"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
