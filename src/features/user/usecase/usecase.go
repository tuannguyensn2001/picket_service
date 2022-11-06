package user_usecase

import (
	"context"
	"myclass_service/src/entities"
)

type IRepository interface {
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	InsertByGoogleAccount(ctx context.Context, entity *entities.User) error
}

type usecase struct {
	repository IRepository
}

func New(repository IRepository) *usecase {
	return &usecase{repository: repository}
}

func (u *usecase) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	return u.repository.FindByEmail(ctx, email)
}

func (u *usecase) CreateByGoogleAccount(ctx context.Context, entity *entities.User) error {
	return u.repository.InsertByGoogleAccount(ctx, entity)
}
