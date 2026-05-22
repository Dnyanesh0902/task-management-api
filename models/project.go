package models

import "time"

type Project struct {
	ID          uint      `gorm:"primaryKey" json:"id"`

	Name        string    `gorm:"size:100;not null" json:"name"`

	Description string    `gorm:"type:text" json:"description"`

	Tasks []Task `gorm:"foreignKey:ProjectID" json:"-"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}