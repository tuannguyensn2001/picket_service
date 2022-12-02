package test_repository

import (
	"context"
	"gorm.io/gorm"
	"picket/src/entities"
	"picket/src/repository"
)

type repo struct {
	repository.Repository
}

func New(db *gorm.DB) *repo {
	return &repo{
		repository.Repository{
			Db: db,
		},
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
