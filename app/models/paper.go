package models

import "time"

type Paper struct {
	ID        string     `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	SubjectID string     `json:"subject_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	TeacherID *string    `json:"teacher_id" gorm:"index;type:uuid" validate:"omitempty,uuid"`
	Name      string     `json:"name" gorm:"not null" validate:"required"`
	Code      string     `json:"code" gorm:"uniqueIndex;not null" validate:"required"`
	IsActive  bool       `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Subject   *Subject   `json:"subject,omitempty" gorm:"foreignKey:SubjectID;references:ID"`
	Teacher   *User      `json:"teacher,omitempty" gorm:"foreignKey:TeacherID;references:ID"`
	Results   []*Result  `json:"results,omitempty" gorm:"foreignKey:PaperID;references:ID"`
}
