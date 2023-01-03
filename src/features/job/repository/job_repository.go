package job_repository

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

func (r *repo) Create(ctx context.Context, job *entities.Job) error {
	db := r.GetDB(ctx).WithContext(ctx)

	m := model{
		Payload:      job.Payload,
		Status:       job.Status,
		ErrorMessage: job.ErrorMessage,
		Topic:        job.Topic,
	}

	if err := db.Create(&m).Error; err != nil {
		return err
	}

	job.Id = m.Id
	job.CreatedAt = m.CreatedAt
	job.UpdatedAt = m.UpdatedAt

	return nil
}

func (r *repo) UpdateSuccess(ctx context.Context, jobId int) error {
	db := r.GetDB(ctx).WithContext(ctx)

	if err := db.Model(&model{}).Where("id = ?", jobId).Update("status", entities.SUCCESS).Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) UpdateFail(ctx context.Context, jobId int, errorMessage string) error {
	db := r.GetDB(ctx).WithContext(ctx)

	if err := db.Model(&model{}).Where("id = ?", jobId).Updates(map[string]interface{}{
		"status":        entities.FAIL,
		"error_message": errorMessage,
	}).Error; err != nil {
		return err
	}

	return nil
}
