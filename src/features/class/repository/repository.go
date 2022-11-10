package class_repository

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

func (r *repo) Create(ctx context.Context, class *entities.Class) error {
	db := r.GetDB(ctx)
	model := model{
		Name:        class.Name,
		Description: class.Description,
	}
	if err := db.WithContext(ctx).Create(&model).Error; err != nil {
		return err
	}
	class.Id = model.Id
	class.CreatedAt = model.CreatedAt
	class.UpdatedAt = model.UpdatedAt
	return nil
}
