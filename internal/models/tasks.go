package models

import "time"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Status string

const (
	StatusIncomplete Status = "Incomplete"
	StatusComplete   Status = "Complete"
	StatusInProgress Status = "In Progress"
)
