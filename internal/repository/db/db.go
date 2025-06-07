package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"todo-api/internal/models"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{db: pool}
}

func (r *Repository) CreateTask(ctx context.Context, task *models.Task) error {
	query := `
        INSERT INTO tasks (title, description, status)
        VALUES ($1, $2, $3)
        RETURNING id, created_at AT TIME ZONE 'UTC', updated_at AT TIME ZONE 'UTC'
    `
	err := r.db.QueryRow(ctx, query, task.Title, task.Description, task.Status).
		Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return fmt.Errorf("error creating task: %w", err)
	}
	return nil
}

func (r *Repository) GetTasks(ctx context.Context) ([]models.Task, error) {
	query := `
        SELECT id, title, description, status, 
               created_at AT TIME ZONE 'UTC' as created_at, 
               updated_at AT TIME ZONE 'UTC' as updated_at
        FROM tasks
        ORDER BY created_at DESC
    `
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying tasks: %w", err)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning task: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over tasks: %w", err)
	}

	return tasks, nil
}

func (r *Repository) UpdateTask(ctx context.Context, task *models.Task) error {
	query := `
        UPDATE tasks
        SET title = $2, description = $3, status = $4, updated_at = NOW()
        WHERE id = $1
        RETURNING created_at AT TIME ZONE 'UTC', updated_at AT TIME ZONE 'UTC'
    `
	err := r.db.QueryRow(ctx, query, task.ID, task.Title, task.Description, task.Status).
		Scan(&task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return fmt.Errorf("error updating task: %w", err)
	}
	return nil
}

func (r *Repository) DeleteTask(ctx context.Context, id int) error {
	query := "DELETE FROM tasks WHERE id = $1"
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting task: %w", err)
	}
	return nil
}
