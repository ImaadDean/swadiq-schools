package models

import "time"

// Attendance represents a student's daily attendance
type Attendance struct {
	ID        string           `json:"id" validate:"required,uuid"`
	StudentID string           `json:"student_id" validate:"required,uuid"`
	ClassID   string           `json:"class_id" validate:"required,uuid"`
	Date      time.Time        `json:"date" validate:"required"`
	Status    AttendanceStatus `json:"status" validate:"required,oneof=present absent late"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt *time.Time       `json:"deleted_at,omitempty"`
	Student   *Student         `json:"student,omitempty"`
	Class     *Class           `json:"class,omitempty"`
}
