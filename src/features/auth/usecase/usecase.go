package auth_usecase

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"myclass_service/src/config"
	"myclass_service/src/entities"
	auth_struct "myclass_service/src/features/auth/struct"
	errpkg "myclass_service/src/packages/err"
)

type IRepository interface {
}

var tracer = otel.Tracer("auth_usecase")

type IOauthService interface {
	GetAccessTokenFromCode(ctx context.Context, code string) (string, error)
	GetUserProfileByAccessToken(ctx context.Context, accessToken string) (*entities.User, error)
}

type IUserUsecase interface {
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	CreateByGoogleAccount(ctx context.Context, entity *entities.User) error
}

type usecase struct {
	repository   IRepository
	oauthService IOauthService
	userUsecase  IUserUsecase
	config       config.IConfig
}

func New(repository IRepository, oauthService IOauthService, userUsecase IUserUsecase, config config.IConfig) *usecase {
	return &usecase{repository: repository, oauthService: oauthService, userUsecase: userUsecase, config: config}
}

func (u *usecase) LoginGoogle(ctx context.Context, code string) (*auth_struct.LoginGoogleOutput, error) {

	ctx, span := tracer.Start(ctx, "login google")
	defer span.End()

	ctx, span = tracer.Start(ctx, "validate code")
	if len(code) == 0 {
		return nil, errpkg.Auth.CodeNotValid
	}
	span.End()

	ctx, span = tracer.Start(ctx, "get access token from code")
	result, err := u.oauthService.GetAccessTokenFromCode(ctx, code)
	span.End()
	if err != nil {
		return nil, err
	}

	googleAccount, err := u.oauthService.GetUserProfileByAccessToken(ctx, result)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}

	user, err := u.userUsecase.GetByEmail(ctx, googleAccount.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = &entities.User{
			Email:    googleAccount.Email,
			Username: googleAccount.Username,
			Profile: &entities.Profile{
				Avatar: googleAccount.Profile.Avatar,
			},
		}
		err = u.userUsecase.CreateByGoogleAccount(ctx, user)
		if err != nil {
			return nil, err
		}
	}

	token, err := u.GenerateToken(ctx, *user)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}

	res := auth_struct.LoginGoogleOutput{AccessToken: token}

	return &res, nil
}
