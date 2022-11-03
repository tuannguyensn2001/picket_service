package routes

import (
	"context"
	"google.golang.org/grpc"
	"myclass_service/src/features/auth_transport"
	authpb "myclass_service/src/pb/auth"
)

func RouteGrpc(ctx context.Context, s *grpc.Server) {
	authTransport := auth_transport.New(ctx)

	authpb.RegisterAuthServiceServer(s, authTransport)
}
