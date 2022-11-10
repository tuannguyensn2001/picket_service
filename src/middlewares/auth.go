package middlewares

import (
	"context"
	"errors"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/metadata"
	"myclass_service/src/config"
	auth_usecase "myclass_service/src/features/auth/usecase"
	user_repository "myclass_service/src/features/user/repository"
	user_usecase "myclass_service/src/features/user/usecase"
	errpkg "myclass_service/src/packages/err"
	"myclass_service/src/routes"
	oauth_service "myclass_service/src/services/oauth"
	"strconv"
)

// exampleAuthFunc is used by a middleware to authenticate requests

var ErrCannotGetMetadata = errors.New("can't get metadata")
var ErrCannotGetPattern = errors.New("can't get pattern")

func Auth(config config.IConfig) func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return ctx, ErrCannotGetMetadata
		}
		pattern := md.Get("pattern")

		if len(pattern) != 0 && !checkIsPrivateRoute(pattern[0]) {
			return ctx, nil
		}

		oauthService := oauth_service.New(config)

		userRepository := user_repository.New(config.GetDB())
		userUsecase := user_usecase.New(userRepository)

		authUsecase := auth_usecase.New(nil, oauthService, userUsecase, config)

		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			panic(errpkg.Auth.Unauthorized)
			return ctx, nil
		}

		userId, err := authUsecase.VerifyToken(ctx, token)
		if err != nil {
			panic(errpkg.Auth.Unauthorized)
			return ctx, nil
		}
		// WARNING: in production define your own type to avoid context collisions
		newCtx := context.WithValue(ctx, "userId", strconv.Itoa(userId))

		return newCtx, nil
	}
}

func checkIsPrivateRoute(pattern string) bool {
	for _, item := range routes.PrivateRoutes {
		if item == pattern {
			return true
		}
	}

	return false
}
