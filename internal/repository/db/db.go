package db

import (
	"context"
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
        RETURNING id, created_at, updated_at
    `
	return r.db.QueryRow(ctx, query, task.Title, task.Description, task.Status).
		Scan(&task.ID, &task.CreatedAt, &task.UpdateAt)
}

func (r *Repository) GetTasks(ctx context.Context) ([]models.Task, error) {
	query := `
        SELECT id, title, description, status, created_at, updated_at
        FROM tasks
        ORDER BY created_at DESC
    `
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
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
			&task.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, rows.Err()
}

func (r *Repository) UpdateTask(ctx context.Context, task *models.Task) error {
	query := `
        UPDATE tasks
        SET title = $2, description = $3, status = $4, updated_at = NOW()
        WHERE id = $1
        RETURNING updated_at
    `
	return r.db.QueryRow(ctx, query, task.ID, task.Title, task.Description, task.Status).
		Scan(&task.UpdateAt)
}

func (r *Repository) DeleteTask(ctx context.Context, id int) error {
	query := "DELETE FROM tasks WHERE id = $1"
	_, err := r.db.Exec(ctx, query, id)
	return err
}
