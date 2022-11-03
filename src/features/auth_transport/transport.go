package auth_transport

import (
	"context"
	"myclass_service/src/packages/err"
	authpb "myclass_service/src/pb/auth"
)

type transport struct {
	authpb.UnimplementedAuthServiceServer
}

func New(ctx context.Context) *transport {
	return &transport{}
}

func (t *transport) Register(ctx context.Context, request *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	return nil, err.General.BadRequest
}
