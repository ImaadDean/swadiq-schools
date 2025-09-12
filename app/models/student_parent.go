package models

import "time"

// StudentParent represents the many-to-many relationship between a student and a parent/guardian.
type StudentParent struct {
	ID           string           `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	StudentID    string           `json:"student_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	ParentID     string           `json:"parent_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	Relationship RelationshipType `json:"relationship" gorm:"not null;type:varchar(20)" validate:"required"`
	IsPrimary    bool             `json:"is_primary" gorm:"default:false"`
	CreatedAt    time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time        `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    *time.Time       `json:"deleted_at,omitempty" gorm:"index"`
	Student      *Student         `json:"student,omitempty" gorm:"foreignKey:StudentID;references:ID"`
	Parent       *Parent          `json:"parent,omitempty" gorm:"foreignKey:ParentID;references:ID"`
}
