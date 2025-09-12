package models

import "time"

type Subject struct {
	ID           string      `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	Name         string      `json:"name" gorm:"not null" validate:"required"`
	Code         string      `json:"code" gorm:"uniqueIndex;not null" validate:"required"`
	DepartmentID *string     `json:"department_id,omitempty" gorm:"index;type:uuid" validate:"omitempty,uuid"`
	IsActive     bool        `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    *time.Time  `json:"deleted_at,omitempty" gorm:"index"`
	Department   *Department `json:"department,omitempty" gorm:"foreignKey:DepartmentID;references:ID"`
	Papers       []*Paper    `json:"papers,omitempty" gorm:"foreignKey:SubjectID;references:ID"` // list of papers
	Classes      []*Class    `json:"classes,omitempty" gorm:"many2many:class_subjects;"`
}
