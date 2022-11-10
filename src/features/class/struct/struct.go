package class_struct

type CreateClassInput struct {
	Name        string `validate:"required"`
	Description string `validate:"required"`
	UserId      int    `validate:"required"`
}
