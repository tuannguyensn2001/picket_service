package test_repository

import (
	"context"
	"picket/src/entities"
)

func (r *repo) FindTestMultipleChoiceAnswer(ctx context.Context, multipleChoiceId int) ([]entities.TestMultipleChoiceAnswer, error) {
	db := r.GetDB(ctx).WithContext(ctx)
	answers := make([]multipleChoiceAnswer, 0)
	if err := db.Where("test_multiple_choice_id", multipleChoiceId).Find(&answers).Error; err != nil {
		return nil, err
	}

	result := make([]entities.TestMultipleChoiceAnswer, len(answers))

	for index, item := range answers {
		result[index] = entities.TestMultipleChoiceAnswer{
			Id:                   item.Id,
			TestMultipleChoiceId: item.TestMultipleChoiceId,
			Answer:               item.Answer,
			Score:                item.Score,
			Type:                 item.Type,
			CreatedAt:            item.CreatedAt,
			UpdatedAt:            item.UpdatedAt,
		}
	}

	return result, nil

}
