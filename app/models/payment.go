package models

import "time"

// Payment represents a payment made by a student for a fee
type Payment struct {
	ID         string     `json:"id" validate:"required,uuid"`
	FeeID      string     `json:"fee_id" validate:"required,uuid"`
	Amount     float64    `json:"amount" validate:"required,gt=0"`
	PaidBy     string     `json:"paid_by" validate:"required,uuid"`
	PaidAt     time.Time  `json:"paid_at" validate:"required"`
	CreatedAt  time.Time  `json:"created_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
	Fee        *Fee       `json:"fee,omitempty"`
	PaidByUser *User      `json:"paid_by_user,omitempty"`
}
