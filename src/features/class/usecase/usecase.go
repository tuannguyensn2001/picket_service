package class_usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"myclass_service/src/entities"
	class_struct "myclass_service/src/features/class/struct"
	errpkg "myclass_service/src/packages/err"
	"myclass_service/src/repository"
)

type IRepository interface {
	Create(ctx context.Context, class *entities.Class) error
	repository.IBaseRepository
}

type usecase struct {
	repository IRepository
}

func New(repository IRepository) *usecase {
	return &usecase{repository: repository}
}

func (u *usecase) Create(ctx context.Context, input class_struct.CreateClassInput) error {
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return errpkg.General.BadRequest
	}

	ctx = u.repository.BeginTransaction(ctx)

	class := entities.Class{
		Name:        input.Name,
		Description: input.Description,
	}
	err = u.repository.Create(ctx, &class)
	if err != nil {
		zap.S().Error(err)
		u.repository.Rollback(ctx)
		return err
	}

	zap.S().Info(class)

	return nil
}
