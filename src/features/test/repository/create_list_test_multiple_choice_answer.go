package test_repository

import (
	"context"
	"myclass_service/src/entities"
)

func (r *repo) CreateListTestMultipleChoiceAnswers(ctx context.Context, list []entities.TestMultipleChoiceAnswer) error {
	db := r.GetDB(ctx)

	model := make([]multipleChoiceAnswer, len(list))

	for index, item := range list {
		model[index] = multipleChoiceAnswer{
			TestMultipleChoiceId: item.TestMultipleChoiceId,
			Answer:               item.Answer,
			Score:                item.Score,
			Type:                 item.Type,
		}
	}

	if err := db.WithContext(ctx).Create(model).Error; err != nil {
		return err
	}

	return nil

}
