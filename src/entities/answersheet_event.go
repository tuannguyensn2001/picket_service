package entities

import "time"

type AnswerSheetEvent struct {
	Id        int        `json:"id"`
	TestId    int        `json:"test_id"`
	UserId    int        `json:"user_id"`
	Event     string     `json:"event"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

const (
	START = "START"
	DOING = "DOING"
	END   = "END"
)
