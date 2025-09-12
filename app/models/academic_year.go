package models

import "time"

// AcademicYear represents an academic year/term in the school
type AcademicYear struct {
	ID        string     `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	Name      string     `json:"name" gorm:"uniqueIndex;not null" validate:"required"`
	StartDate time.Time  `json:"start_date" gorm:"not null;index" validate:"required"`
	EndDate   time.Time  `json:"end_date" gorm:"not null;index" validate:"required"`
	IsCurrent bool       `json:"is_current" gorm:"default:false;index"`
	IsActive  bool       `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}

// Term represents a term/semester within an academic year
type Term struct {
	ID             string        `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	AcademicYearID string        `json:"academic_year_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	Name           string        `json:"name" gorm:"not null" validate:"required"`
	StartDate      time.Time     `json:"start_date" gorm:"not null;index" validate:"required"`
	EndDate        time.Time     `json:"end_date" gorm:"not null;index" validate:"required"`
	IsCurrent      bool          `json:"is_current" gorm:"default:false;index"`
	IsActive       bool          `json:"is_active" gorm:"default:true"`
	CreatedAt      time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      *time.Time    `json:"deleted_at,omitempty" gorm:"index"`
	AcademicYear   *AcademicYear `json:"academic_year,omitempty" gorm:"foreignKey:AcademicYearID;references:ID"`
}
