package repository

import (
	"context"
	"todo-api/internal/models"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task *models.Task) error
	GetTasks(ctx context.Context) ([]models.Task, error)
	UpdateTask(ctx context.Context, task *models.Task) error
	DeleteTask(ctx context.Context, id int) error
}
