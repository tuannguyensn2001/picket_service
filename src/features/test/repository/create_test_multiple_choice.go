package test_repository

import (
	"context"
	"myclass_service/src/entities"
)

func (r *repo) CreateTestMultipleChoice(ctx context.Context, test *entities.TestMultipleChoice) error {
	db := r.GetDB(ctx)
	model := multipleChoice{
		FilePath: test.FilePath,
		Score:    test.Score,
	}
	if err := db.WithContext(ctx).Create(&model).Error; err != nil {
		return err
	}
	test.Id = model.Id
	test.CreatedAt = model.CreatedAt
	test.UpdatedAt = model.UpdatedAt

	return nil
}
