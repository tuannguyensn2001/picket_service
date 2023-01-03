package job_repository

import "time"

type model struct {
	Id           int        `gorm:"column:id"`
	Payload      string     `gorm:"column:payload"`
	Status       string     `gorm:"column:status"`
	ErrorMessage string     `gorm:"column:error_message"`
	Topic        string     `gorm:"column:topic"`
	CreatedAt    *time.Time `gorm:"column:created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at"`
}

func (model) TableName() string {
	return "jobs"
}
