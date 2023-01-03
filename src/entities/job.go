package entities

import "time"

type Job struct {
	Id           int        `json:"id"`
	Payload      string     `json:"payload"`
	Status       string     `json:"status"`
	ErrorMessage string     `json:"error_message"`
	Topic        string     `json:"topic"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

const (
	INIT    = "INIT"
	SUCCESS = "SUCCESS"
	FAIL    = "FAIL"
)
