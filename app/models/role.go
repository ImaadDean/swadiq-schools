package models

import "time"

// Role represents a user role (e.g., admin, bursar)
type Role struct {
	ID          string        `json:"id" validate:"required,uuid"`
	Name        string        `json:"name" validate:"required"`
	IsActive    bool          `json:"is_active"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	DeletedAt   *time.Time    `json:"deleted_at,omitempty"`
	Permissions []*Permission `json:"permissions,omitempty"` // optional for JSON responses
}

// Permission represents a fine-grained action a role can perform
type Permission struct {
	ID        string     `json:"id" validate:"required,uuid"`
	Name      string     `json:"name" validate:"required"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// RolePermission is a join table linking roles and permissions
type RolePermission struct {
	ID           string     `json:"id" validate:"required,uuid"`
	RoleID       string     `json:"role_id" validate:"required,uuid"`
	PermissionID string     `json:"permission_id" validate:"required,uuid"`
	CreatedAt    time.Time  `json:"created_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}
