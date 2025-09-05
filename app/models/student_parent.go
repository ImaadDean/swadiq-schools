package models

import "time"

// StudentParent represents the many-to-many relationship between a student and a parent/guardian.
type StudentParent struct {
	ID           string           `json:"id" validate:"required,uuid"`
	StudentID    string           `json:"student_id" validate:"required,uuid"`
	ParentID     string           `json:"parent_id" validate:"required,uuid"`
	Relationship RelationshipType `json:"relationship" validate:"required"`
	IsPrimary    bool             `json:"is_primary"`
	CreatedAt    time.Time        `json:"created_at"`
	DeletedAt    *time.Time       `json:"deleted_at,omitempty"`
	Student      *Student         `json:"student,omitempty"`
	Parent       *Parent          `json:"parent,omitempty"`
}
