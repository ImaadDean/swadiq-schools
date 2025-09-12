package models

import "time"

type Parent struct {
	ID            string           `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	FirstName     string           `json:"first_name" gorm:"not null" validate:"required"`
	LastName      string           `json:"last_name" gorm:"not null" validate:"required"`
	Email         *string          `json:"email,omitempty" gorm:"index" validate:"omitempty,email"`
	Phone         *string          `json:"phone,omitempty" gorm:"index"`
	Address       *string          `json:"address,omitempty" gorm:"type:text"`
	IsActive      bool             `json:"is_active" gorm:"default:true"`
	CreatedAt     time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time        `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt     *time.Time       `json:"deleted_at,omitempty" gorm:"index"`
	Students      []*Student       `json:"students,omitempty" gorm:"many2many:student_parents;"`
	StudentParent []*StudentParent `json:"student_parent,omitempty"`
}
