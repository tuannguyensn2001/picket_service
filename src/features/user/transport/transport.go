package user_transport

import (
	"context"
	"picket/src/entities"
	userpb "picket/src/pb/user"
	"picket/src/utils"
)

type IUsecase interface {
	GetById(ctx context.Context, id int) (*entities.User, error)
}

type transport struct {
	userpb.UnimplementedUserServiceServer
	usecase IUsecase
}

func New(ctx context.Context, usecase IUsecase) *transport {
	return &transport{usecase: usecase}
}

func (t *transport) GetProfile(ctx context.Context, request *userpb.GetProfileRequest) (*userpb.GetProfileResponse, error) {
	userId, err := utils.GetAuth(ctx)
	if err != nil {
		panic(err)
	}

	user, err := t.usecase.GetById(ctx, userId)
	if err != nil {
		panic(err)
	}

	resp := &userpb.GetProfileResponse{
		Message: "success",
		Data: &userpb.User{
			Email:    user.Email,
			Username: user.Username,
			Profile: &userpb.Profile{
				Avatar: user.Profile.Avatar,
			},
		},
	}
	return resp, nil
}
