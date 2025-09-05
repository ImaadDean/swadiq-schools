package models

import "time"

type Subject struct {
	ID        string     `json:"id" validate:"required,uuid"`
	Name      string     `json:"name" validate:"required"`
	Code      string     `json:"code" validate:"required"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Papers    []*Paper   `json:"papers,omitempty"` // list of papers
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Classes   []*Class   `json:"classes,omitempty"`
}
