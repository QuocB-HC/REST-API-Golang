package models

import "time"

type TodoItem struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Image       string     `json:"image"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
