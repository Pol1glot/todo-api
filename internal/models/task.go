package models

import "time"

type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdateAt    time.Time
}
