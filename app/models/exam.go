package models

import "time"

// Exam represents an exam event for a class
type Exam struct {
	ID             string        `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	Name           string        `json:"name" gorm:"not null" validate:"required"`
	ClassID        string        `json:"class_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	AcademicYearID *string       `json:"academic_year_id,omitempty" gorm:"index;type:uuid" validate:"omitempty,uuid"`
	TermID         *string       `json:"term_id,omitempty" gorm:"index;type:uuid" validate:"omitempty,uuid"`
	StartDate      time.Time     `json:"start_date" gorm:"not null;index" validate:"required"`
	EndDate        time.Time     `json:"end_date" gorm:"not null;index" validate:"required"`
	IsActive       bool          `json:"is_active" gorm:"default:true"`
	CreatedAt      time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      *time.Time    `json:"deleted_at,omitempty" gorm:"index"`
	Class          *Class        `json:"class,omitempty" gorm:"foreignKey:ClassID;references:ID"`
	AcademicYear   *AcademicYear `json:"academic_year,omitempty" gorm:"foreignKey:AcademicYearID;references:ID"`
	Term           *Term         `json:"term,omitempty" gorm:"foreignKey:TermID;references:ID"`
	Results        []*Result     `json:"results,omitempty" gorm:"foreignKey:ExamID;references:ID"`
}
