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
