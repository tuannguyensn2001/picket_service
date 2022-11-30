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

type multipleChoiceAnswer struct {
	Id                   int        `gorm:"column:id"`
	TestMultipleChoiceId int        `gorm:"column:test_multiple_choice_id"`
	Answer               string     `gorm:column:answer"`
	Score                float64    `gorm:"column:score"`
	Type                 int        `gorm:"column:type"`
	CreatedAt            *time.Time `gorm:"column:created_at"`
	UpdatedAt            *time.Time `gorm:"column:updated_at"`
}

func (multipleChoiceAnswer) TableName() string {
	return "test_multiple_choice_answers"
}

type multipleChoice struct {
	Id        int        `gorm:"column:id"`
	FilePath  string     `gorm:"column:file_path"`
	Score     float64    `gorm:"column:score"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (multipleChoice) TableName() string {
	return "test_multiple_choice"
}

type content struct {
	Id         int        `gorm:"column:id"`
	TestId     int        `gorm:"column:test_id"`
	TypeableId int        `gorm:"column:typeable_id"`
	Typeable   int        `gorm:"column:typeable"`
	CreatedAt  *time.Time `gorm:"column:created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at"`
}

func (content) TableName() string {
	return "test_content"
}
