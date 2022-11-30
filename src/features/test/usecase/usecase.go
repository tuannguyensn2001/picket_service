package test_usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"myclass_service/src/entities"
	test_struct "myclass_service/src/features/test/struct"
	errpkg "myclass_service/src/packages/err"
	"myclass_service/src/repository"
	"myclass_service/src/utils"
	"time"
)

type IRepository interface {
	Create(ctx context.Context, test *entities.Test) error
	CreateTestContent(ctx context.Context, test *entities.TestContent) error
	CreateTestMultipleChoice(ctx context.Context, test *entities.TestMultipleChoice) error
	CreateListTestMultipleChoiceAnswers(ctx context.Context, list []entities.TestMultipleChoiceAnswer) error
	repository.IBaseRepository
	FindByTestId(ctx context.Context, id int) (*entities.Test, error)
	FindContentByTestId(ctx context.Context, testId int) (*entities.TestContent, error)
}

type usecase struct {
	repository IRepository
}

func New(repository IRepository) *usecase {
	return &usecase{repository: repository}
}

func (u *usecase) Create(ctx context.Context, input test_struct.CreateTestInput, userId int) error {

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		zap.S().Error(err)
		return errpkg.General.BadRequest
	}

	timeStart, err := u.ParseTimeTest(ctx, input.TimeStart)
	if err != nil {
		zap.S().Error(err)
		return errpkg.General.BadRequest
	}
	timeEnd, err := u.ParseTimeTest(ctx, input.TimeEnd)
	if err != nil {
		zap.S().Error(err)
		return errpkg.General.BadRequest
	}

	if timeStart != nil && timeEnd != nil {
		if timeStart.After(*timeEnd) {
			zap.S().Error("time start before time end")
			return errpkg.General.BadRequest
		}
	}

	test := entities.Test{
		Code:               uuid.New().String(),
		Name:               input.Name,
		TimeToDo:           input.TimeToDo,
		TimeStart:          timeStart,
		TimeEnd:            timeEnd,
		DoOnce:             input.DoOnce,
		Password:           input.Password,
		PreventCheat:       input.PreventCheat,
		IsAuthenticateUser: input.IsAuthenticateUser,
		ShowAnswer:         input.ShowAnswer,
		ShowMark:           input.ShowMark,
		CreatedBy:          userId,
		Version:            1,
	}
	err = u.repository.Create(ctx, &test)
	if err != nil {
		return err
	}

	return nil
}

func (u *usecase) ParseTimeTest(ctx context.Context, val string) (*time.Time, error) {
	if len(val) == 0 {
		return nil, nil
	}
	return utils.ParseTime("HH:MM:SS DD/MM/YYYY", val)
}
