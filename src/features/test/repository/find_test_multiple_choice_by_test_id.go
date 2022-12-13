package test_repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"picket/src/entities"
	"strings"
	"time"
)

func (r *repo) FindTestMultipleChoiceByTestId(ctx context.Context, testId int) (*entities.TestMultipleChoice, error) {
	result, err := r.FindTestMultipleChoiceFromRedisByTestId(ctx, testId)
	if err == nil && result != nil {
		return result, nil
	}
	result, err = r.FindTestMultipleChoiceFromDBByTestId(ctx, testId)
	if err != nil {
		return nil, err
	}
	go r.SaveTestMultipleChoiceToRedis(ctx, testId, result)
	return result, nil
}

func (r *repo) FindTestMultipleChoiceFromDBByTestId(ctx context.Context, testId int) (*entities.TestMultipleChoice, error) {
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

func (r *repo) FindTestMultipleChoiceFromRedisByTestId(ctx context.Context, testId int) (*entities.TestMultipleChoice, error) {
	status := r.redis.Get(ctx, fmt.Sprintf("test_%d_multiple_choice", testId))
	if status.Err() == nil {
		var result entities.TestMultipleChoice
		err := json.NewDecoder(strings.NewReader(status.Val())).Decode(&result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, status.Err()
}

func (r *repo) SaveTestMultipleChoiceToRedis(ctx context.Context, testId int, multipleChoice *entities.TestMultipleChoice) error {
	b := &bytes.Buffer{}
	err := json.NewEncoder(b).Encode(multipleChoice)
	if err != nil {
		zap.S().Error(err)
		return err
	}
	status := r.redis.Set(ctx, fmt.Sprintf("test_%d_multiple_choice", testId), b.String(), 1*time.Hour)
	if status.Err() != nil {
		zap.S().Error(err)
		return status.Err()
	}
	return nil
}
