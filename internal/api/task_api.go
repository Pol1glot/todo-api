package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"todo-api/internal/dtos"
)

type TaskAPI interface {
	CreateTask(c *fiber.Ctx) error
	GetTasks(c *fiber.Ctx) error
	UpdateTask(c *fiber.Ctx) error
	DeleteTask(c *fiber.Ctx) error
}

type TaskService interface {
	CreateTask(ctx context.Context, task *dtos.CreateTaskDTO) (*dtos.TaskDTO, error)
	GetTasks(ctx context.Context) ([]dtos.TaskDTO, error)
	UpdateTask(ctx context.Context, id int, task *dtos.UpdateTaskDTO) (*dtos.TaskDTO, error)
	DeleteTask(ctx context.Context, id int) error
}
