package models

import "time"

type Student struct {
	ID            string           `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	StudentID     string           `json:"student_id" gorm:"uniqueIndex;not null" validate:"required"`
	FirstName     string           `json:"first_name" gorm:"not null" validate:"required"`
	LastName      string           `json:"last_name" gorm:"not null" validate:"required"`
	DateOfBirth   *time.Time       `json:"date_of_birth,omitempty" gorm:"type:date"`
	Gender        *Gender          `json:"gender,omitempty" gorm:"type:varchar(10)" validate:"omitempty,oneof=male female other"`
	Address       *string          `json:"address,omitempty" gorm:"type:text"`
	ClassID       *string          `json:"class_id,omitempty" gorm:"index;type:uuid" validate:"omitempty,uuid"`
	IsActive      bool             `json:"is_active" gorm:"default:true"`
	CreatedAt     time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time        `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt     *time.Time       `json:"deleted_at,omitempty" gorm:"index"`
	Class         *Class           `json:"class,omitempty" gorm:"foreignKey:ClassID;references:ID"` // optional
	Parents       []*Parent        `json:"parents,omitempty" gorm:"many2many:student_parents;"`
	StudentParent []*StudentParent `json:"student_parent,omitempty"`
}
