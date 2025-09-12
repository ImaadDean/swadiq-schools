package models

import "time"

// Expense represents a school expense
type Expense struct {
	ID         string     `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	CategoryID string     `json:"category_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	Title      string     `json:"title" gorm:"not null" validate:"required"`
	Amount     float64    `json:"amount" gorm:"not null;type:decimal(10,2)" validate:"required,gt=0"`
	Currency   string     `json:"currency" gorm:"not null;default:'USD';type:varchar(3)" validate:"required,len=3"`
	Date       time.Time  `json:"date" gorm:"not null;index;type:date" validate:"required"`
	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Category   *Category  `json:"category,omitempty" gorm:"foreignKey:CategoryID;references:ID"` // optional for JSON responses
}
