package auth_repository

import (
	"gorm.io/gorm"
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
