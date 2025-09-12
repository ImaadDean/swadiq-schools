package models

import "time"

type User struct {
	ID        string     `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	Email     string     `json:"email" gorm:"uniqueIndex;not null" validate:"required,email"`
	Password  string     `json:"-" gorm:"not null" validate:"required,min=8"`
	FirstName string     `json:"first_name" gorm:"not null" validate:"required"`
	LastName  string     `json:"last_name" gorm:"not null" validate:"required"`
	Phone     string     `json:"phone,omitempty" gorm:"type:varchar(20)"`
	IsActive  bool       `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Roles     []*Role    `json:"roles,omitempty" gorm:"many2many:user_roles;"` // Many-to-many relationship
}

type Session struct {
	ID        string     `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	UserID    string     `json:"user_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	ExpiresAt time.Time  `json:"expires_at" gorm:"not null;index"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	User      *User      `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID"`
}
