package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type TodoItem struct {
	ID          uuid.UUID      `json:"id" gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Title       string         `json:"title" gorm:"column:title;type:varchar(100);not null"`
	Image       datatypes.JSON `json:"image" gorm:"column:image"`
	Description string         `json:"description" gorm:"column:description"`
	Status      string         `json:"status" gorm:"column:status;type:status_enum;default:'Doing';not null"`
	CreatedAt   time.Time      `json:"created_at" gorm:"column:created_at;type:timestamp"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"column:updated_at;type:timestamp"`
}

func (TodoItem) TableName() string { return "todo_items" }

type TodoItemCreation struct {
	Title       string `json:"title" binding:"required" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *string `json:"status" gorm:"column:status"`
	UpdatedAt   time.Time
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }
