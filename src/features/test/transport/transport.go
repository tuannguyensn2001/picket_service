package test_transport

import (
	"context"
	"picket/src/entities"
	test_struct "picket/src/features/test/struct"
	testpb "picket/src/pb/test"
	"picket/src/utils"
)

type IUsecase interface {
	Create(ctx context.Context, input test_struct.CreateTestInput, userId int) error
	CreateContent(ctx context.Context, input test_struct.CreateTestContentInput) error
	GetTestsByUserId(ctx context.Context, userId int) ([]entities.Test, error)
	GetPreview(ctx context.Context, id int) (*entities.Test, error)
}

type transport struct {
	usecase IUsecase
	testpb.UnimplementedTestServiceServer
}

func New(ctx context.Context, usecase IUsecase) *transport {
	return &transport{usecase: usecase}
}

func (t *transport) Create(ctx context.Context, request *testpb.CreateTestRequest) (*testpb.CreateTestResponse, error) {

	input := test_struct.CreateTestInput{
		Name:               request.Name,
		TimeToDo:           int(request.TimeToDo),
		TimeStart:          request.TimeStart,
		TimeEnd:            request.TimeEnd,
		DoOnce:             request.DoOnce,
		Password:           request.Password,
		PreventCheat:       uint8(request.PreventCheat),
		IsAuthenticateUser: request.IsAuthenticateUser,
		ShowMark:           uint8(request.ShowMark),
		ShowAnswer:         uint8(request.ShowAnswer),
	}
	userId, err := utils.GetAuth(ctx)
	if err != nil {
		panic(err)
	}
	err = t.usecase.Create(ctx, input, userId)
	if err != nil {
		panic(err)
	}

	resp := testpb.CreateTestResponse{
		Message: "success",
	}
	return &resp, nil
}
