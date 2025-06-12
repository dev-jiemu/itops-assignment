package models

import "time"

const (
	ISSUE_STATUS_PENDING     = "PENDING"
	ISSUE_STATUS_IN_PROGRESS = "IN_PROGRESS"
	ISSUE_STATUS_COMPLETED   = "COMPLETED"
	ISSUE_STATUS_CANCELLED   = "CANCELLED"
)

type Issue struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	User        *User     `json:"user,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
