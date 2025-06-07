package dtos

import "time"

// TaskDTO представляет собой DTO для задачи
type TaskDTO struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTaskDTO struct {
	Title       string
	Description string
}

// UpdateTaskDTO представляет собой DTO для обновления существующей задачи
type UpdateTaskDTO struct {
	Title       string
	Description string
	Status      string
}
