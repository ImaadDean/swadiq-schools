package models

import "time"

// Category represents a category of expense
type Category struct {
	ID        string     `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	Name      string     `json:"name" gorm:"uniqueIndex;not null" validate:"required"`
	IsActive  bool       `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Expenses  []*Expense `json:"expenses,omitempty" gorm:"foreignKey:CategoryID;references:ID"`
}
