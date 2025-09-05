package models

import "time"

type User struct {
	ID        string     `json:"id" validate:"required,uuid"`
	Email     string     `json:"email" validate:"required,email"`
	Password  string     `json:"-" validate:"required,min=8"`
	FirstName string     `json:"first_name" validate:"required"`
	LastName  string     `json:"last_name" validate:"required"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Roles     []*Role    `json:"roles,omitempty"` // Optional for JSON responses
}

type Session struct {
	ID        string
	UserID    int
	ExpiresAt time.Time
	CreatedAt time.Time
}
