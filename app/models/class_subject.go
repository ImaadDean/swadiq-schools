package models

import "time"

// ClassSubject represents the relationship between a class and a subject, forming a many-to-many join table.
type ClassSubject struct {
	ID        string     `json:"id" validate:"required,uuid"`
	ClassID   string     `json:"class_id" validate:"required,uuid"`
	SubjectID string     `json:"subject_id" validate:"required,uuid"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}