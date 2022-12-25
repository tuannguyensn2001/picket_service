package job_usecase

import (
	"context"
	"picket/src/entities"
)

type IRepository interface {
	Create(ctx context.Context, job *entities.Job) error
	UpdateSuccess(ctx context.Context, jobId int) error
	UpdateFail(ctx context.Context, jobId int, errorMessage string) error
}

type usecase struct {
	repository IRepository
}

func New(repository IRepository) *usecase {
	return &usecase{repository: repository}
}

func (u *usecase) Create(ctx context.Context, job *entities.Job) error {
	err := u.repository.Create(ctx, job)
	if err != nil {
		return err
	}

	return nil
}

func (u *usecase) UpdateSuccess(ctx context.Context, jobId int) error {
	return u.repository.UpdateSuccess(ctx, jobId)
}

func (u *usecase) UpdateFail(ctx context.Context, jobId int, errFail error) error {
	return u.repository.UpdateFail(ctx, jobId, errFail.Error())
}
