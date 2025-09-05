package models

import "time"

type Class struct {
	ID        string     `json:"id" validate:"required,uuid"`
	Name      string     `json:"name" validate:"required"`
	TeacherID *string    `json:"teacher_id,omitempty" validate:"omitempty,uuid"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Subjects  []*Subject `json:"subjects,omitempty"`
}
