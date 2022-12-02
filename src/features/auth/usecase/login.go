package auth_usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	auth_struct "picket/src/features/auth/struct"
	errpkg "picket/src/packages/err"
)

func (u *usecase) Login(ctx context.Context, input auth_struct.LoginInput) (*auth_struct.LoginOutput, error) {
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		zap.S().Error(err)
		return nil, errpkg.Auth.LoginFailed
	}

	user, err := u.userUsecase.GetByEmail(ctx, input.Email)
	if err != nil {
		zap.S().Error(err)
		return nil, errpkg.Auth.LoginFailed
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		zap.S().Error(err)
		return nil, errpkg.Auth.LoginFailed
	}

	token, err := u.GenerateToken(ctx, *user)
	if err != nil {
		zap.S().Error(err)
		return nil, errpkg.Auth.LoginFailed
	}

	result := auth_struct.LoginOutput{
		AccessToken: token,
	}

	return &result, nil
}
