package models

import "time"

type Paper struct {
	ID        string     `json:"id" validate:"required,uuid"`
	SubjectID string     `json:"subject_id" validate:"required,uuid"`
	TeacherID *string    `json:"teacher_id" validate:"omitempty,uuid"`
	Name      string     `json:"name" validate:"required"`
	Code      string     `json:"code" validate:"required"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Subject   *Subject   `json:"subject,omitempty"`
	Teacher   *User      `json:"teacher,omitempty"`
}
