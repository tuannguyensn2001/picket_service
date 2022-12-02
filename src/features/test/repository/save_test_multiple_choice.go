package test_repository

import (
	"context"
	"picket/src/entities"
)

func (r *repo) SaveTestMultipleChoice(ctx context.Context, entity *entities.TestMultipleChoice) error {
	model := multipleChoice{
		Id:        entity.Id,
		FilePath:  entity.FilePath,
		Score:     entity.Score,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
	db := r.GetDB(ctx)
	if err := db.WithContext(ctx).Save(&model).Error; err != nil {
		return err
	}
	return nil
}
