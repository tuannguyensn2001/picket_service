package entities

import (
	"time"
)

type Test struct {
	Id                 int        `json:"id"`
	Code               string     `json:"code"`
	Name               string     `json:"name"`
	TimeToDo           int        `json:"time_to_do"`
	TimeStart          *time.Time `json:"time_start"`
	TimeEnd            *time.Time `json:"time_end"`
	DoOnce             bool       `json:"do_once"`
	Password           string     `json:"password"`
	PreventCheat       uint8      `json:"prevent_cheat"`
	IsAuthenticateUser bool       `json:"is_authenticate_user"`
	ShowMark           uint8      `json:"show_mark"`
	ShowAnswer         uint8      `json:"show_answer"`
	CreatedBy          int        `json:"created_by"`
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at"`
	Version            int        `json:"version"`
}

type TestContent struct {
	Id         int        `json:"id"`
	TestId     int        `json:"test_id"`
	TypeableId int        `json:"typeable_id"`
	Typeable   int        `json:"typeable"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

type TestMultipleChoiceAnswer struct {
	Id                   int        `json:"id"`
	TestMultipleChoiceId int        `json:"test_multiple_choice_id"`
	Answer               string     `json:"answer"`
	Score                float64    `json:"score"`
	Type                 int        `json:"type"`
	CreatedAt            *time.Time `json:"created_at"`
	UpdatedAt            *time.Time `json:"updated_at"`
}

type TestMultipleChoice struct {
	Id        int        `json:"id"`
	FilePath  string     `json:"file_path"`
	Score     float64    `json:"score"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
