package test_struct

type CreateTestInput struct {
	Name               string `validate:"required"`
	TimeToDo           int    `validate:"required,gt=0"`
	TimeStart          string
	TimeEnd            string
	DoOnce             bool
	Password           string
	PreventCheat       uint8 `validate:"required"`
	IsAuthenticateUser bool
	ShowMark           uint8 `validate:"required"`
	ShowAnswer         uint8 `validate:"required"`
}
