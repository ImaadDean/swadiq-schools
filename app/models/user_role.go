package models

import "time"

type UserRole struct {
	ID        string     `json:"id" validate:"required,uuid"`
	UserID    string     `json:"user_id" validate:"required,uuid"`
	RoleID    string     `json:"role_id" validate:"required,uuid"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	User      *User      `json:"user,omitempty"`
	Role      *Role      `json:"role,omitempty"`
}
