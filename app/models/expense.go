package models

import "time"

// Expense represents a school expense
type Expense struct {
	ID         string     `json:"id" validate:"required,uuid"`
	CategoryID string     `json:"category_id" validate:"required,uuid"`
	Title      string     `json:"title" validate:"required"`
	Amount     float64    `json:"amount" validate:"required,gt=0"`
	Date       time.Time  `json:"date" validate:"required"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
	Category   *Category  `json:"category,omitempty"` // optional for JSON responses
}
