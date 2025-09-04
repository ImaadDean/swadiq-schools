package models

import "time"

type Student struct {
	ID          string    `json:"id"`
	StudentID   string    `json:"student_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth *string   `json:"date_of_birth"`
	Gender      *string   `json:"gender"`
	Address     *string   `json:"address"`
	ParentID    *string   `json:"parent_id"`
	ClassID     *int      `json:"class_id"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Parent      *Parent   `json:"parent,omitempty"`
}
