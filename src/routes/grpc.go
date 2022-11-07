package routes

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"myclass_service/src/config"
	auth_transport "myclass_service/src/features/auth/transport"
	auth_usecase "myclass_service/src/features/auth/usecase"
	media_transport "myclass_service/src/features/media/transport"
	media_usecase "myclass_service/src/features/media/usecase"
	user_repository "myclass_service/src/features/user/repository"
	user_transport "myclass_service/src/features/user/transport"
	user_usecase "myclass_service/src/features/user/usecase"
	authpb "myclass_service/src/pb/auth"
	userpb "myclass_service/src/pb/user"
	oauth_service "myclass_service/src/services/oauth"
	storage_service "myclass_service/src/services/storage"
	"net/http"
)

type handler = func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

func RouteGrpc(ctx context.Context, s *grpc.Server, config config.IConfig) {
	oauthService := oauth_service.New(config)

	userRepository := user_repository.New(config.GetDB())
	userUsecase := user_usecase.New(userRepository)
	userTransport := user_transport.New(ctx, userUsecase)

	authUsecase := auth_usecase.New(nil, oauthService, userUsecase, config)
	authTransport := auth_transport.New(ctx, authUsecase)

	authpb.RegisterAuthServiceServer(s, authTransport)
	userpb.RegisterUserServiceServer(s, userTransport)
}

func RouteGw(ctx context.Context, gw *runtime.ServeMux, conn *grpc.ClientConn) {
	lists := []handler{authpb.RegisterAuthServiceHandler, userpb.RegisterUserServiceHandler}

	for _, item := range lists {
		err := item(ctx, gw, conn)
		if err != nil {
			zap.S().Fatalln(err)
		}
	}

	storageService := storage_service.New("fs", "src/storage")
	mediaUsecase := media_usecase.New(storageService)
	mediaTransport := media_transport.New(ctx, mediaUsecase)

	gw.HandlePath(http.MethodPost, "/api/v1/media/upload", mediaTransport.Upload)
}
