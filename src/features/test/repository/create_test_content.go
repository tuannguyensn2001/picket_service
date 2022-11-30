package test_repository

import (
	"context"
	"myclass_service/src/entities"
)

func (r *repo) CreateTestContent(ctx context.Context, test *entities.TestContent) error {
	db := r.GetDB(ctx)
	model := content{
		TestId:     test.TestId,
		Typeable:   test.Typeable,
		TypeableId: test.TypeableId,
	}
	if err := db.WithContext(ctx).Create(&model).Error; err != nil {
		return err
	}

	test.Id = model.Id
	test.CreatedAt = model.CreatedAt
	test.UpdatedAt = model.UpdatedAt

	return nil
}
