package user_transport

import (
	"context"
	"go.uber.org/zap"
	userpb "myclass_service/src/pb/user"
	"myclass_service/src/utils"
)

type transport struct {
	userpb.UnimplementedUserServiceServer
}

func New(ctx context.Context) *transport {
	return &transport{}
}

func (t *transport) GetProfile(ctx context.Context, request *userpb.GetProfileRequest) (*userpb.GetProfileResponse, error) {
	userId, err := utils.GetAuth(ctx)
	zap.S().Info(userId, err)
	resp := &userpb.GetProfileResponse{
		Message: "success",
	}
	return resp, nil
}
