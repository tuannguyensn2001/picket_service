package user_repository

import (
	"context"
	"gorm.io/gorm"
	"myclass_service/src/entities"
	"myclass_service/src/repository"
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

func (r *repo) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	db := r.GetDB(ctx)
	var model user
	if err := db.WithContext(ctx).Where("email = ?", email).First(&model).Error; err != nil {
		return nil, err
	}

	result := entities.User{
		Email:    model.Email,
		Username: model.Username,
	}
	return &result, nil
}

func (r *repo) InsertByGoogleAccount(ctx context.Context, entity *entities.User) error {
	db := r.GetDB(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		model := user{
			Username: entity.Username,
			Email:    entity.Email,
			Status:   active,
			Type:     type_account_google,
		}

		if err := tx.WithContext(ctx).Create(&model).Error; err != nil {
			return err
		}

		p := profile{
			UserId: model.Id,
			Avatar: entity.Profile.Avatar,
		}

		if err := tx.WithContext(ctx).Create(&p).Error; err != nil {
			return err
		}

		entity.Type = model.Type
		entity.Status = model.Status

		return nil
	})
}
