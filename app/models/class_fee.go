package models

import "time"

type ClassFee struct {
	ID         string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ClassID    string    `json:"class_id" gorm:"not null;index;type:uuid"`
	FeeTypeID  string    `json:"fee_type_id" gorm:"not null;index;type:uuid"`
	TermID     string    `json:"term_id" gorm:"not null;index;type:uuid"`
	Amount     float64   `json:"amount" gorm:"not null;type:decimal(10,2)"`
	IsRequired bool      `json:"is_required" gorm:"not null;default:true"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Term       *Term     `json:"term,omitempty" gorm:"foreignKey:TermID;references:ID"`
	FeeType    *FeeType  `json:"fee_type,omitempty" gorm:"foreignKey:FeeTypeID;references:ID"`
}
