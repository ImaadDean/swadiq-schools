package models

import "time"

type UserRole struct {
	ID        string     `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	UserID    string     `json:"user_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	RoleID    string     `json:"role_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	User      *User      `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID"`
	Role      *Role      `json:"role,omitempty" gorm:"foreignKey:RoleID;references:ID"`
}
