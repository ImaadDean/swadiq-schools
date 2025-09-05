package models

import "time"

// Fee represents a fee assigned to a student
type Fee struct {
	ID        string     `json:"id" validate:"required,uuid"`
	StudentID string     `json:"student_id" validate:"required,uuid"`
	Title     string     `json:"title" validate:"required"`
	Amount    float64    `json:"amount" validate:"required,gt=0"`
	Paid      bool       `json:"paid"`
	DueDate   time.Time  `json:"due_date" validate:"required"`
	PaidAt    *time.Time `json:"paid_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Student   *Student   `json:"student,omitempty"`
}
