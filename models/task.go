package models

import "time"

type Task struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Title string `gorm:"size=100;not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Status string `gorm:"default:pending" json:"status"`
	Priority string `gorm:"default:normal" json:"priority"`
	ProjectID uint `json:"project_id"`
	Project Project `gorm:"foreignKey:ProjectID"`
	AssignedToID uint `json:"assigned_to_id"`
	AssignedUser User `gorm:"foreignKey:AssignedToID"`
	DueDate time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

}