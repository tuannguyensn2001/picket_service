package auth_transport

import (
	"context"
	auth_struct "myclass_service/src/features/auth/struct"
	authpb "myclass_service/src/pb/auth"
)

type transport struct {
	authpb.UnimplementedAuthServiceServer
	usecase IUsecase
}

type IUsecase interface {
	LoginGoogle(ctx context.Context, code string) (*auth_struct.LoginGoogleOutput, error)
	Register(ctx context.Context, input auth_struct.RegisterInput) error
	Login(ctx context.Context, input auth_struct.LoginInput) (*auth_struct.LoginOutput, error)
}

func New(ctx context.Context, usecase IUsecase) *transport {
	return &transport{usecase: usecase}
}

func (t *transport) Register(ctx context.Context, request *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	input := auth_struct.RegisterInput{
		Email:    request.Email,
		Password: request.Password,
		Username: request.Username,
	}
	err := t.usecase.Register(ctx, input)
	if err != nil {
		panic(err)
	}
	return &authpb.RegisterResponse{
		Message: "success",
	}, nil
}

func (t *transport) LoginGoogle(ctx context.Context, request *authpb.LoginGoogleRequest) (*authpb.LoginGoogleResponse, error) {
	result, err := t.usecase.LoginGoogle(ctx, request.GetCode())
	if err != nil {
		panic(err)
	}
	return &authpb.LoginGoogleResponse{Message: "success", Data: &authpb.LoginGoogleOutput{AccessToken: result.AccessToken}}, nil
}

func (t *transport) Login(ctx context.Context, request *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	result, err := t.usecase.Login(ctx, auth_struct.LoginInput{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		panic(err)
	}

	resp := authpb.LoginResponse{
		Message: "success",
		Data: &authpb.LoginOutput{
			AccessToken: result.AccessToken,
		},
	}
	return &resp, nil
}
