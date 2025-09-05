package models

import "time"

// Grade represents a grading rule, e.g., A, B, C
type Grade struct {
	ID        string     `json:"id" validate:"required,uuid"`
	Name      string     `json:"name" validate:"required"`
	MinMarks  float64    `json:"min_marks" validate:"gte=0"`
	MaxMarks  float64    `json:"max_marks" validate:"gte=0"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
