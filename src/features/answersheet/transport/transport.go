package answersheet_transport

import (
	"context"
	answersheet_struct "picket/src/features/answersheet/struct"
)

type IUsecase interface {
	Start(ctx context.Context, testId int, userId int) (*answersheet_struct.StartOutput,error)
}

type transport struct {
	usecase IUsecase
}

func New(ctx context.Context, usecase IUsecase) *transport  {
	return &transport{usecase: usecase}
}
