package class_usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"myclass_service/src/entities"
	class_struct "myclass_service/src/features/class/struct"
	errpkg "myclass_service/src/packages/err"
	"myclass_service/src/packages/slices"
	"myclass_service/src/repository"
	"myclass_service/src/utils"
	"strings"
)

const (
	randomLength = 5
)

type IRepository interface {
	Create(ctx context.Context, class *entities.Class) error
	repository.IBaseRepository
	AddTeacherToClass(ctx context.Context, userId int, classId int) error
	GetAllCodes(ctx context.Context) ([]string, error)
	FindByQuery(ctx context.Context, params class_struct.QueryClass) ([]entities.Class, error)
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

	codes, err := u.repository.GetAllCodes(ctx)
	if err != nil {
		zap.S().Error(err)
		return err
	}

	code := utils.RandomWithLength(randomLength)
	for {
		if !slices.Includes(codes, code) {
			break
		}
		code = utils.RandomWithLength(5)
	}

	class := entities.Class{
		Name:        input.Name,
		Description: input.Description,
		Code:        strings.ToUpper(code),
	}
	ctx = u.repository.BeginTransaction(ctx)

	err = u.repository.Create(ctx, &class)
	if err != nil {
		zap.S().Error(err)
		u.repository.Rollback(ctx)
		return err
	}
	err = u.repository.AddTeacherToClass(ctx, input.UserId, class.Id)
	if err != nil {
		zap.S().Error(err)
		u.repository.Rollback(ctx)
		return err
	}

	u.repository.Commit(ctx)

	return nil
}
