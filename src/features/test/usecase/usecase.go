package test_usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"picket/src/entities"
	test_struct "picket/src/features/test/struct"
	errpkg "picket/src/packages/err"
	randompkg "picket/src/packages/random"
	"picket/src/repository"
	"picket/src/utils"
	"strings"
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
	FindTestMultipleChoiceByTestId(ctx context.Context, testId int) (*entities.TestMultipleChoice, error)
	FindTestByUserId(ctx context.Context, userId int) ([]entities.Test, error)
	FindTestMultipleChoiceAnswer(ctx context.Context, multipleChoiceId int) ([]entities.TestMultipleChoiceAnswer, error)
	FindByCode(ctx context.Context, code string) (*entities.Test, error)
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
		Code:               strings.ToUpper(randompkg.StringWithLength(5)),
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
