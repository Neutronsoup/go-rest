package models

import (
	"time"
)

type Task struct {
	ID          uint      `json:"id"  gorm:"promaryKey"`
	Title       string    `json:"title" gorm:"size:200"`
	Description string    `json:"description" gorm:"type:text"`
	Done        bool      `json:"done" gorm:"default:false"`
	OwnerID     uint      `json:"owner_id" gorm:"index"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}
type CreateTaskDTO struct {
	Title       string `json:"title" binding:"required,min=3,max=200"`
	Description string `json:"description"`
}
type UpdateTaskDTO struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Done        *bool   `json:"done,omitempty"`
}
