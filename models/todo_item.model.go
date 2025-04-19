package models

import "time"

type TodoItem struct {
	ID          string     `json:"id" gorm:"column:id"`
	Title       string     `json:"title" gorm:"column:title"`
	Image       string     `json:"image" gorm:"column:image"`
	Description string     `json:"description" gorm:"column:description"`
	Status      string     `json:"status" gorm:"column:status"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (TodoItem) TableName() string { return "todo_items" }

type TodoItemCreation struct {
	Title       string `json:"title" binding:"required" gorm:"column:title"`
	Image       string `json:"image" gorm:"column:image"`
	Description string `json:"description" gorm:"column:description"`
	// Status      string     `json:"status" gorm:"column:status;type:status_enum;default:'Doing'"`
}

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title"`
	Image       string  `json:"image" gorm:"column:image"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *string `json:"status" gorm:"column:status"`
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }
