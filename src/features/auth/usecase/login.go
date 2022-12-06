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
	ctx, span := tracer.Start(ctx, "start login")
	defer span.End()

	ctx, span = tracer.Start(ctx, "validate")
	validate := validator.New()
	err := validate.Struct(input)
	span.End()
	if err != nil {
		zap.S().Error(err)
		return nil, errpkg.Auth.LoginFailed
	}

	ctx, span = tracer.Start(ctx, "get user by email")
	user, err := u.userUsecase.GetByEmail(ctx, input.Email)
	span.End()
	if err != nil {
		zap.S().Error(err)
		return nil, errpkg.Auth.LoginFailed
	}

	ctx, span = tracer.Start(ctx, "compare password")
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	span.End()
	if err != nil {
		zap.S().Error(err)
		return nil, errpkg.Auth.LoginFailed
	}

	ctx, span = tracer.Start(ctx, "generat jwt")
	token, err := u.GenerateToken(ctx, *user)
	span.End()
	if err != nil {
		zap.S().Error(err)
		return nil, errpkg.Auth.LoginFailed
	}

	result := auth_struct.LoginOutput{
		AccessToken: token,
	}

	return &result, nil
}
