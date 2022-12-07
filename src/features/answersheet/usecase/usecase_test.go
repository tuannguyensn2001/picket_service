package answersheet_usecase

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"picket/src/entities"
	errpkg "picket/src/packages/err"
	"testing"
	"time"
)

func TestUsecase_Start(t *testing.T) {
	errpkg.LoadErrorFromPath("/Users/tuannguyen/Workspace/Go/picket_service/error.yml")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := NewMockIRepository(ctrl)
	testUsecase := NewMockITestUsecase(ctrl)

	usecase := New(repository, testUsecase)

	t.Run("return error when get test fail", func(t *testing.T) {
		ctx := context.TODO()
		testUsecase.EXPECT().GetById(ctx, gomock.Any()).Return(nil, gorm.ErrRecordNotFound)
		_, err := usecase.Start(ctx, 1, 1)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	})

	t.Run("return error when test haven't started", func(t *testing.T) {
		ctx := context.TODO()
		timeStart := time.Now().Add(1 * time.Hour)
		testUsecase.EXPECT().GetById(ctx, gomock.Any()).Return(&entities.Test{
			TimeStart: &timeStart,
		}, nil)
		_, err := usecase.Start(ctx, 1, 1)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errpkg.Answersheet.TimeNotValid)
	})

	t.Run("return error when test have done", func(t *testing.T) {
		ctx := context.TODO()
		timeEnd := time.Now().Add(-1 * time.Hour)
		testUsecase.EXPECT().GetById(ctx, gomock.Any()).Return(&entities.Test{
			TimeEnd: &timeEnd,
		}, nil)
		_, err := usecase.Start(ctx, 1, 1)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errpkg.Answersheet.TimeNotValid)
	})
}

func TestCheckUserDoingTest(t *testing.T) {
	errpkg.LoadErrorFromPath("/Users/tuannguyen/Workspace/Go/picket_service/error.yml")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := NewMockIRepository(ctrl)
	testUsecase := NewMockITestUsecase(ctrl)
	usecase := New(repository, testUsecase)

	t.Run("user don't have event ", func(t *testing.T) {
		ctx := context.TODO()
		repository.EXPECT().GetLatestEvent(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return([]entities.AnswerSheetEvent{}, nil)
		check, err := usecase.CheckUserDoingTest(ctx, 1, 1)
		assert.False(t, check)
		assert.Nil(t, err)
	})

	t.Run("user has 1 event and it's start", func(t *testing.T) {
		ctx := context.TODO()
		repository.EXPECT().GetLatestEvent(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return([]entities.AnswerSheetEvent{
			{
				Event: entities.START,
			},
		}, nil)
		check, err := usecase.CheckUserDoingTest(ctx, 1, 1)
		assert.True(t, check)
		assert.Nil(t, err)
	})
	t.Run("user has 1 event and it's doing", func(t *testing.T) {
		ctx := context.TODO()
		repository.EXPECT().GetLatestEvent(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return([]entities.AnswerSheetEvent{
			{
				Event: entities.DOING,
			},
		}, nil)
		check, err := usecase.CheckUserDoingTest(ctx, 1, 1)
		assert.True(t, check)
		assert.Nil(t, err)
	})

	t.Run("user has more than 1 event", func(t *testing.T) {
		ctx := context.TODO()
		repository.EXPECT().GetLatestEvent(ctx,gomock.Any(),gomock.Any(),gomock.Any()).Return([]entities.AnswerSheetEvent{
			{
				Event: entities.END,
			},
			{
				Event: entities.START,
			},
		},nil)
		check,err := usecase.CheckUserDoingTest(ctx,1,1)
		assert.False(t, check)
		assert.Nil(t, err)
	})
	t.Run("user has more than 1 event and false", func(t *testing.T) {
		ctx := context.TODO()
		repository.EXPECT().GetLatestEvent(ctx,gomock.Any(),gomock.Any(),gomock.Any()).Return([]entities.AnswerSheetEvent{
			{
				Event: entities.DOING,
			},
			{
				Event: entities.START,
			},
		},nil)
		check,err := usecase.CheckUserDoingTest(ctx,1,1)
		assert.True(t, check)
		assert.Nil(t, err)
	})

}
