package auth_transport

import (
	"context"
	auth_struct "myclass_service/src/features/auth/struct"
	"myclass_service/src/packages/err"
	authpb "myclass_service/src/pb/auth"
)

type transport struct {
	authpb.UnimplementedAuthServiceServer
	usecase IUsecase
}

type IUsecase interface {
	LoginGoogle(ctx context.Context, code string) (*auth_struct.LoginGoogleOutput, error)
}

func New(ctx context.Context, usecase IUsecase) *transport {
	return &transport{usecase: usecase}
}

func (t *transport) Register(ctx context.Context, request *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	panic(errpkg.General.BadRequest)
	return nil, errpkg.General.BadRequest
}

func (t *transport) LoginGoogle(ctx context.Context, request *authpb.LoginGoogleRequest) (*authpb.LoginGoogleResponse, error) {
	_, err := t.usecase.LoginGoogle(ctx, request.GetCode())
	if err != nil {
		panic(err)
	}
	return &authpb.LoginGoogleResponse{Message: "success"}, nil
}
