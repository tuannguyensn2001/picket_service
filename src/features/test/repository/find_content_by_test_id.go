package test_repository

import (
	"context"
	"picket/src/entities"
)

func (r *repo) FindContentByTestId(ctx context.Context, testId int) (*entities.TestContent, error) {
	db := r.GetDB(ctx)
	var model content
	if err := db.WithContext(ctx).Where("test_id = ?", testId).Order("id desc").First(&model).Error; err != nil {
		return nil, err
	}

	content := entities.TestContent{
		TestId:     model.TestId,
		Id:         model.Id,
		Typeable:   model.Typeable,
		TypeableId: model.TypeableId,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}

	return &content, nil

}
