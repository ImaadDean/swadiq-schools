package models

import "time"

type Student struct {
	ID            string           `json:"id" validate:"required,uuid"`
	StudentID     string           `json:"student_id" validate:"required"`
	FirstName     string           `json:"first_name" validate:"required"`
	LastName      string           `json:"last_name" validate:"required"`
	DateOfBirth   *string          `json:"date_of_birth,omitempty"`
	Gender        *Gender          `json:"gender,omitempty" validate:"omitempty,oneof=male female other"`
	Address       *string          `json:"address,omitempty"`
	ClassID       *string          `json:"class_id,omitempty" validate:"omitempty,uuid"`
	IsActive      bool             `json:"is_active"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
	Class         *Class           `json:"class,omitempty"` // optional
	DeletedAt     *time.Time       `json:"deleted_at,omitempty"`
	Parents       []*Parent        `json:"parents,omitempty"`
	StudentParent []*StudentParent `json:"student_parent,omitempty"`
}
