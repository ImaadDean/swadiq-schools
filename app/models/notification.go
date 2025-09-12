package models

import "time"

// Notification represents an email sent to a user with full history and attachments
type Notification struct {
	ID             string        `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()" validate:"required,uuid"`
	Subject        string        `json:"subject" gorm:"not null" validate:"required"`
	Body           string        `json:"body" gorm:"not null;type:text" validate:"required"`
	RecipientID    string        `json:"recipient_id" gorm:"not null;index;type:uuid" validate:"required,uuid"`
	RecipientType  RecipientType `json:"recipient_type" gorm:"not null;index;type:varchar(10)" validate:"required,oneof=student parent teacher"`
	Email          string        `json:"email" gorm:"not null;index" validate:"required,email"`
	IsSent         bool          `json:"is_sent" gorm:"default:false;index"`
	SentAt         *time.Time    `json:"sent_at,omitempty" gorm:"index"`
	Template       string        `json:"template" gorm:"type:varchar(100)"`
	RetryCount     int           `json:"retry_count" gorm:"default:0" validate:"gte=0"`
	AttachmentURLs []string      `json:"attachment_urls" gorm:"type:text[]"`
	CreatedAt      time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      *time.Time    `json:"deleted_at,omitempty" gorm:"index"`
}
