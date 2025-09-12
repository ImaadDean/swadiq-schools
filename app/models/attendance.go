package models

import "time"

// Attendance represents a student's daily attendance
type Attendance struct {
	ID        string           `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	StudentID string           `json:"student_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	ClassID   string           `json:"class_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	Date      time.Time        `json:"date" gorm:"not null;index;type:date" validate:"required"`
	Status    AttendanceStatus `json:"status" gorm:"not null;type:varchar(10)" validate:"required,oneof=present absent late"`
	CreatedAt time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time        `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time       `json:"deleted_at,omitempty" gorm:"index"`
	Student   *Student         `json:"student,omitempty" gorm:"foreignKey:StudentID;references:ID"`
	Class     *Class           `json:"class,omitempty" gorm:"foreignKey:ClassID;references:ID"`
}
