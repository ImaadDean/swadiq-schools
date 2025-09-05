package models

import "time"

// Result stores a student's marks for a paper in an exam
type Result struct {
	ID        string     `json:"id" validate:"required,uuid"`
	ExamID    string     `json:"exam_id" validate:"required,uuid"`
	StudentID string     `json:"student_id" validate:"required,uuid"`
	PaperID   string     `json:"paper_id" validate:"required,uuid"`
	Marks     float64    `json:"marks" validate:"gte=0"`
	GradeID   *string    `json:"grade_id,omitempty" validate:"omitempty,uuid"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Grade     *Grade     `json:"grade,omitempty"` // optional for JSON responses
	Exam      *Exam      `json:"exam,omitempty"`
	Student   *Student   `json:"student,omitempty"`
	Paper     *Paper     `json:"paper,omitempty"`
}
