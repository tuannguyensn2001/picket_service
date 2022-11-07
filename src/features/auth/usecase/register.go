package auth_usecase

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"myclass_service/src/entities"
	auth_struct "myclass_service/src/features/auth/struct"
	errpkg "myclass_service/src/packages/err"
)

func (u *usecase) Register(ctx context.Context, input auth_struct.RegisterInput) error {
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		zap.S().Error(err)
		return errpkg.General.BadRequest
	}

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), 5)
	if err != nil {
		return err
	}
	user, err := u.userUsecase.GetByEmail(ctx, input.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if user != nil {
		return errpkg.Auth.AccountExisted
	}

	entity := entities.User{
		Email:    input.Email,
		Password: string(password),
		Username: input.Username,
	}

	err = u.userUsecase.CreateAccount(ctx, &entity)
	if err != nil {
		return err
	}

	return nil
}
