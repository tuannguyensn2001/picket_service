package entities

import "time"

type Class struct {
	Id          int
	Name        string
	Description string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
