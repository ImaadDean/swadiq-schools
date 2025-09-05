package models

import "time"

type Parent struct {
	ID            string           `json:"id" validate:"required,uuid"`
	FirstName     string           `json:"first_name" validate:"required"`
	LastName      string           `json:"last_name" validate:"required"`
	Email         *string          `json:"email,omitempty" validate:"omitempty,email"`
	Phone         *string          `json:"phone,omitempty"`
	Address       *string          `json:"address,omitempty"`
	IsActive      bool             `json:"is_active"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
	DeletedAt     *time.Time       `json:"deleted_at,omitempty"`
	Students      []*Student       `json:"students,omitempty"`
	StudentParent []*StudentParent `json:"student_parent,omitempty"`
}
