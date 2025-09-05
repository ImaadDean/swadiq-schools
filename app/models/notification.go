package models

import "time"

// Notification represents an email sent to a user with full history and attachments
type Notification struct {
	ID             string        `json:"id" validate:"required,uuid"`
	Subject        string        `json:"subject" validate:"required"`
	Body           string        `json:"body" validate:"required"`
	RecipientID    string        `json:"recipient_id" validate:"required,uuid"`
	RecipientType  RecipientType `json:"recipient_type" validate:"required,oneof=student parent teacher"`
	Email          string        `json:"email" validate:"required,email"`
	IsSent         bool          `json:"is_sent"`
	SentAt         *time.Time    `json:"sent_at,omitempty"`
	Template       string        `json:"template"`
	RetryCount     int           `json:"retry_count" validate:"gte=0"`
	AttachmentURLs []string      `json:"attachment_urls"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	DeletedAt      *time.Time    `json:"deleted_at,omitempty"`
}
