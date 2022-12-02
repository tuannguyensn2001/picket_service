package test_repository

import (
	"context"
	"picket/src/entities"
)

func (r *repo) FindTestMultipleChoiceByTestId(ctx context.Context, testId int) (*entities.TestMultipleChoice, error) {
	var model multipleChoice

	query := `select tmc.* from test_multiple_choice tmc join test_content tc on tmc.id = tc.typeable_id where tc.typeable = ? and tc.test_id = ? order by id desc limit 1`
	db := r.GetDB(ctx)

	if err := db.WithContext(ctx).Raw(query, entities.MULTIPLE_CHOICE, testId).First(&model).Error; err != nil {
		return nil, err
	}

	result := entities.TestMultipleChoice{
		Id:        model.Id,
		FilePath:  model.FilePath,
		Score:     model.Score,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}

	return &result, nil
}
