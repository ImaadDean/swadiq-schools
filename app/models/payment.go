package models

import "time"

// Payment represents a payment made by a student for a fee
type Payment struct {
	ID            string        `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	FeeID         string        `json:"fee_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	Amount        float64       `json:"amount" gorm:"not null;type:decimal(10,2)" validate:"required,gt=0"`
	PaidBy        string        `json:"paid_by" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	PaymentMethod string        `json:"payment_method" gorm:"type:varchar(50)"`
	TransactionID *string       `json:"transaction_id,omitempty" gorm:"index"`
	Status        PaymentStatus `json:"status" gorm:"not null;default:'pending';index;type:varchar(20)" validate:"required"`
	PaidAt        time.Time     `json:"paid_at" gorm:"not null;index" validate:"required"`
	CreatedAt     time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt     *time.Time    `json:"deleted_at,omitempty" gorm:"index"`
	Fee           *Fee          `json:"fee,omitempty" gorm:"foreignKey:FeeID;references:ID"`
	PaidByUser    *User         `json:"paid_by_user,omitempty" gorm:"foreignKey:PaidBy;references:ID"`
}
