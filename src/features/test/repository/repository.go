package test_repository

import (
	"context"
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
	"picket/src/entities"
	"picket/src/repository"
	"sync"
)

type repo struct {
	repository.Repository
	redis               *redis.Client
	s                   sync.Once
	findTestById        sync.Mutex
	findContentByTestId sync.Mutex
}

func New(db *gorm.DB, redis *redis.Client) *repo {
	return &repo{
		Repository: repository.Repository{
			Db: db,
		},
		redis: redis,
	}
}

func (r *repo) Create(ctx context.Context, test *entities.Test) error {
	model := model{
		Code:               test.Code,
		Name:               test.Name,
		TimeToDo:           test.TimeToDo,
		TimeStart:          test.TimeStart,
		TimeEnd:            test.TimeEnd,
		DoOnce:             test.DoOnce,
		Password:           test.Password,
		PreventCheat:       test.PreventCheat,
		IsAuthenticateUser: test.IsAuthenticateUser,
		ShowMark:           test.ShowMark,
		ShowAnswer:         test.ShowAnswer,
		CreatedBy:          test.CreatedBy,
		Version:            test.Version,
	}

	db := r.GetDB(ctx)
	if err := db.WithContext(ctx).Create(&model).Error; err != nil {
		return err
	}
	test.Id = model.Id
	test.CreatedAt = model.CreatedAt
	test.UpdatedAt = model.UpdatedAt

	return nil

}
