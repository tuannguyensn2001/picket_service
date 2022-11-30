package test_repository

import (
	"gorm.io/gorm"
	"time"
)

type model struct {
	Id                 int        `gorm:"column:id"`
	Code               string     `gorm:"column:code"`
	Name               string     `gorm:"column:name"`
	TimeToDo           int        `gorm:"column:time_to_do"`
	TimeStart          *time.Time `gorm:"column:time_start"`
	TimeEnd            *time.Time `gorm:"column:time_end"`
	DoOnce             bool       `gorm:"column:do_once"`
	Password           string     `gorm:"column:password"`
	PreventCheat       uint8      `gorm:"column:prevent_cheat`
	IsAuthenticateUser bool       `gorm:"column:is_authenticate_user"`
	ShowMark           uint8      `gorm:"column:show_mark"`
	ShowAnswer         uint8      `gorm:"column:show_answer"`
	Version            int        `gorm:"column:version"`
	CreatedBy          int        `gorm:"column:created_by"`
	CreatedAt          *time.Time `gorm:"column:created_at"`
	UpdatedAt          *time.Time `gorm:"column:updated_at"`
	DeletedAt          gorm.DeletedAt
}

func (model) TableName() string {
	return "tests"
}
