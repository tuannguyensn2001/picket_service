package routes

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	auth_transport "myclass_service/src/features/auth/transport"
	media_transport "myclass_service/src/features/media/transport"
	authpb "myclass_service/src/pb/auth"
	"net/http"
)

type handler = func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

func RouteGrpc(ctx context.Context, s *grpc.Server) {
	authTransport := auth_transport.New(ctx)

	authpb.RegisterAuthServiceServer(s, authTransport)
}

func RouteGw(ctx context.Context, gw *runtime.ServeMux, conn *grpc.ClientConn) {
	lists := []handler{authpb.RegisterAuthServiceHandler}

	for _, item := range lists {
		err := item(ctx, gw, conn)
		if err != nil {
			zap.S().Fatalln(err)
		}
	}

	mediaTransport := media_transport.New(ctx)

	gw.HandlePath(http.MethodPost, "/api/v1/media/upload", mediaTransport.Upload)
}
