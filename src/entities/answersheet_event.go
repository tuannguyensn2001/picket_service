package entities

import "time"

type AnswerSheetEvent struct {
	Id     int
	TestId int
	UserId int
	Event string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

const (
	START = "START"
	DOING = "DOING"
	END = "END"
)