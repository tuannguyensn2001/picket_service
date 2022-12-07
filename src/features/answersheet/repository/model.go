package answersheet_repository

import "time"

type event struct {
	Id int `gorm:"column:id"`
	UserId int `gorm:"column:user_id"`
	TestId int `gorm:"column:test_id"`
	Event string `gorm:"column:event"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (event) TableName() string  {
	return "answersheet_event"
}