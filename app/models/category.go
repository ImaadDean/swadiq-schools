package models

import "time"

// Category represents a category of expense
type Category struct {
	ID        string     `json:"id" validate:"required,uuid"`
	Name      string     `json:"name" validate:"required"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
