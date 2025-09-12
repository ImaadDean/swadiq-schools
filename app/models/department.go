package models

import "time"

type Department struct {
	ID                   string     `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	Name                 string     `json:"name" gorm:"not null" validate:"required"`
	Code                 string     `json:"code" gorm:"uniqueIndex;not null" validate:"required"`
	Description          *string    `json:"description,omitempty" gorm:"type:text"`
	HeadOfDepartmentID   *string    `json:"head_of_department_id,omitempty" gorm:"index;type:uuid" validate:"omitempty,uuid"`
	AssistantHeadID      *string    `json:"assistant_head_id,omitempty" gorm:"index;type:uuid" validate:"omitempty,uuid"`
	IsActive             bool       `json:"is_active" gorm:"default:true"`
	CreatedAt            time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt            time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt            *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	HeadOfDepartment     *User      `json:"head_of_department,omitempty" gorm:"foreignKey:HeadOfDepartmentID;references:ID"`
	AssistantHead        *User      `json:"assistant_head,omitempty" gorm:"foreignKey:AssistantHeadID;references:ID"`
	Subjects             []*Subject `json:"subjects,omitempty" gorm:"foreignKey:DepartmentID;references:ID"`
}
