package models

import "time"

// Role represents a user role (e.g., admin, bursar)
type Role struct {
	ID          string        `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	Name        string        `json:"name" gorm:"uniqueIndex;not null" validate:"required"`
	IsActive    bool          `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *time.Time    `json:"deleted_at,omitempty" gorm:"index"`
	Permissions []*Permission `json:"permissions,omitempty" gorm:"many2many:role_permissions;"` // optional for JSON responses
	Users       []*User       `json:"users,omitempty" gorm:"many2many:user_roles;"`
}

// Permission represents a fine-grained action a role can perform
type Permission struct {
	ID        string     `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	Name      string     `json:"name" gorm:"uniqueIndex;not null" validate:"required"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Roles     []*Role    `json:"roles,omitempty" gorm:"many2many:role_permissions;"`
}

// RolePermission is a join table linking roles and permissions
type RolePermission struct {
	ID           string     `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	RoleID       string     `json:"role_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	PermissionID string     `json:"permission_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Role         *Role      `json:"role,omitempty" gorm:"foreignKey:RoleID;references:ID"`
	Permission   *Permission `json:"permission,omitempty" gorm:"foreignKey:PermissionID;references:ID"`
}
