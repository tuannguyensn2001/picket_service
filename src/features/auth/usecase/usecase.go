package auth_usecase

import (
	"context"
	"go.uber.org/zap"
	auth_struct "myclass_service/src/features/auth/struct"
	errpkg "myclass_service/src/packages/err"
)

type IRepository interface {
}

type IOauthService interface {
	GetAccessTokenFromCode(ctx context.Context, code string) (string, error)
}

type usecase struct {
	repository   IRepository
	oauthService IOauthService
}

func New(repository IRepository, oauthService IOauthService) *usecase {
	return &usecase{repository: repository, oauthService: oauthService}
}

func (u *usecase) LoginGoogle(ctx context.Context, code string) (*auth_struct.LoginGoogleOutput, error) {
	if len(code) == 0 {
		return nil, errpkg.Auth.CodeNotValid
	}

	result, err := u.oauthService.GetAccessTokenFromCode(ctx, code)
	if err != nil {
		return nil, err
	}
	zap.S().Info(result, err)

	return nil, nil
}
