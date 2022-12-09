package answersheet_repository

import (
	"context"
	"gorm.io/gorm"
	"picket/src/entities"
	repository2 "picket/src/repository"
)

type repository struct {
	repository2.Repository
}

func New(db *gorm.DB) *repository {
	return &repository{
		repository2.Repository{Db: db},
	}
}

func (r *repository) GetLatestEvent(ctx context.Context, userId int, testId int, number int) ([]entities.AnswerSheetEvent, error) {
	events := make([]event, 0)
	db := r.GetDB(ctx).WithContext(ctx)
	if err := db.Where("user_id = ?", userId).Where("test_id = ?", testId).Order("id desc").Limit(number).Find(&events).Error; err != nil {
		return nil, err
	}
	result := make([]entities.AnswerSheetEvent, len(events))

	for index, item := range events {
		result[index] = entities.AnswerSheetEvent{
			Id:        item.Id,
			TestId:    item.TestId,
			UserId:    item.UserId,
			Event:     item.Event,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
	}

	return result, nil
}

func (r *repository) SendEvent(ctx context.Context, model *entities.AnswerSheetEvent) error {
	e := event{
		UserId: model.UserId,
		TestId: model.TestId,
		Event:  model.Event,
	}
	db := r.GetDB(ctx).WithContext(ctx)
	if err := db.Create(&e).Error; err != nil {
		return err
	}
	model.Id = e.Id
	model.CreatedAt = e.CreatedAt
	model.UpdatedAt = e.UpdatedAt
	return nil
}
