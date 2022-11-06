package auth_usecase

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"myclass_service/src/entities"
	auth_struct "myclass_service/src/features/auth/struct"
	errpkg "myclass_service/src/packages/err"
)

type IRepository interface {
}

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
}

func New(repository IRepository, oauthService IOauthService, userUsecase IUserUsecase) *usecase {
	return &usecase{repository: repository, oauthService: oauthService, userUsecase: userUsecase}
}

func (u *usecase) LoginGoogle(ctx context.Context, code string) (*auth_struct.LoginGoogleOutput, error) {
	if len(code) == 0 {
		return nil, errpkg.Auth.CodeNotValid
	}

	result, err := u.oauthService.GetAccessTokenFromCode(ctx, code)
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

	zap.S().Info(user)

	return nil, nil
}
