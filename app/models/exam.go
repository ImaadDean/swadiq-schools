package models

import "time"

// Exam represents an exam event for a class
type Exam struct {
	ID        string     `json:"id" validate:"required,uuid"`
	Name      string     `json:"name" validate:"required"`
	ClassID   string     `json:"class_id" validate:"required,uuid"`
	StartDate time.Time  `json:"start_date" validate:"required"`
	EndDate   time.Time  `json:"end_date" validate:"required"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Class     *Class     `json:"class,omitempty"`
}
