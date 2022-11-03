package auth_transport

import (
	"context"
	"errors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	authpb "myclass_service/src/pb/auth"
)

type transport struct {
	authpb.UnimplementedAuthServiceServer
}

func New(ctx context.Context) *transport {
	return &transport{}
}

func (t *transport) Register(ctx context.Context, request *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	return nil, runtime.HTTPStatusError{
		HTTPStatus: 400,
		Err:        errors.New("abc"),
	}.Err
}
