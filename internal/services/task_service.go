package services

import (
	"context"
	"todo-api/internal/dtos"
	"todo-api/internal/models"
	"todo-api/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, createDTO *dtos.CreateTaskDTO) (*dtos.TaskDTO, error) {
	task := &models.Task{
		Title:       createDTO.Title,
		Description: createDTO.Description,
		Status:      "new", // По умолчанию статус "new"
	}

	err := s.repo.CreateTask(ctx, task)
	if err != nil {
		return nil, err
	}

	return &dtos.TaskDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}

func (s *TaskService) GetTasks(ctx context.Context) ([]dtos.TaskDTO, error) {
	tasks, err := s.repo.GetTasks(ctx)
	if err != nil {
		return nil, err
	}

	taskDTOs := make([]dtos.TaskDTO, len(tasks))
	for i, task := range tasks {
		taskDTOs[i] = dtos.TaskDTO{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		}
	}

	return taskDTOs, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, id int, updateDTO *dtos.UpdateTaskDTO) (*dtos.TaskDTO, error) {
	task := &models.Task{
		ID:          id,
		Title:       updateDTO.Title,
		Description: updateDTO.Description,
		Status:      updateDTO.Status,
	}

	err := s.repo.UpdateTask(ctx, task)
	if err != nil {
		return nil, err
	}

	return &dtos.TaskDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id int) error {
	return s.repo.DeleteTask(ctx, id)
}
