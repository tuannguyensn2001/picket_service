package class_transport

import (
	"context"
	class_struct "picket/src/features/class/struct"
	classpb "picket/src/pb/class"
	"picket/src/utils"
)

type IUsecase interface {
	Create(ctx context.Context, input class_struct.CreateClassInput) error
}

type transport struct {
	classpb.UnimplementedClassServiceServer
	usecase IUsecase
}

func New(ctx context.Context, usecase IUsecase) *transport {
	return &transport{usecase: usecase}
}

func (t *transport) Create(ctx context.Context, request *classpb.CreateClassRequest) (*classpb.CreateClassResponse, error) {
	userId, err := utils.GetAuth(ctx)
	if err != nil {
		panic(err)
	}
	input := class_struct.CreateClassInput{
		UserId:      userId,
		Name:        request.Name,
		Description: request.Description,
	}
	err = t.usecase.Create(ctx, input)
	if err != nil {
		panic(err)
	}

	resp := classpb.CreateClassResponse{
		Message: "success",
	}
	return &resp, nil
}
