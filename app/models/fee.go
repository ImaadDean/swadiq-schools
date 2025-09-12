package models

import "time"

// Fee represents a fee assigned to a student
type Fee struct {
	ID             string        `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	StudentID      string        `json:"student_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	FeeTypeID      string        `json:"fee_type_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	AcademicYearID *string       `json:"academic_year_id,omitempty" gorm:"index;type:uuid" validate:"omitempty,uuid"`
	TermID         *string       `json:"term_id,omitempty" gorm:"index;type:uuid" validate:"omitempty,uuid"`
	Title          string        `json:"title" gorm:"not null" validate:"required"`
	Amount         float64       `json:"amount" gorm:"not null;type:decimal(10,2)" validate:"required,gt=0"`
	Currency       string        `json:"currency" gorm:"not null;default:'USD';type:varchar(3)" validate:"required,len=3"`
	Paid           bool          `json:"paid" gorm:"default:false;index"`
	DueDate        time.Time     `json:"due_date" gorm:"not null;index" validate:"required"`
	PaidAt         *time.Time    `json:"paid_at,omitempty"`
	CreatedAt      time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      *time.Time    `json:"deleted_at,omitempty" gorm:"index"`
	Student        *Student      `json:"student,omitempty" gorm:"foreignKey:StudentID;references:ID"`
	FeeType        *FeeType      `json:"fee_type,omitempty" gorm:"foreignKey:FeeTypeID;references:ID"`
	AcademicYear   *AcademicYear `json:"academic_year,omitempty" gorm:"foreignKey:AcademicYearID;references:ID"`
	Term           *Term         `json:"term,omitempty" gorm:"foreignKey:TermID;references:ID"`
	Payments       []*Payment    `json:"payments,omitempty" gorm:"foreignKey:FeeID;references:ID"`
}
