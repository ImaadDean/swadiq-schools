package models

import "time"

// Schedule represents a class schedule for a subject and teacher
type Schedule struct {
	ID        string    `json:"id" validate:"required,uuid"`
	ClassID   string    `json:"class_id" validate:"required,uuid"`
	SubjectID string    `json:"subject_id" validate:"required,uuid"`
	TeacherID string    `json:"teacher_id" validate:"required,uuid"`
	Day       DayOfWeek `json:"day" validate:"required,oneof=Monday Tuesday Wednesday Thursday Friday Saturday Sunday"`
	StartTime time.Time `json:"start_time" validate:"required"`
	EndTime   time.Time `json:"end_time" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Class     *Class    `json:"class,omitempty"`
	Subject   *Subject  `json:"subject,omitempty"`
	Teacher   *User     `json:"teacher,omitempty"`
}
