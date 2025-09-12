package models

import "time"

// Schedule represents a class schedule for a subject and teacher
type Schedule struct {
	ID        string     `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	ClassID   string     `json:"class_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	SubjectID string     `json:"subject_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	TeacherID string     `json:"teacher_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	Day       DayOfWeek  `json:"day" gorm:"not null;index;type:varchar(10)" validate:"required,oneof=monday tuesday wednesday thursday friday saturday sunday"`
	StartTime string     `json:"start_time" gorm:"not null;type:time" validate:"required"`
	EndTime   string     `json:"end_time" gorm:"not null;type:time" validate:"required"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Class     *Class     `json:"class,omitempty" gorm:"foreignKey:ClassID;references:ID"`
	Subject   *Subject   `json:"subject,omitempty" gorm:"foreignKey:SubjectID;references:ID"`
	Teacher   *User      `json:"teacher,omitempty" gorm:"foreignKey:TeacherID;references:ID"`
}
