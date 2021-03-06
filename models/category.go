package models

import "time"

// Category stores product category
type Category struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name" form:"name" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
