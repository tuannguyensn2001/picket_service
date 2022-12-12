package test_repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"picket/src/entities"
	"strings"
	"time"
)

func (r *repo) FindContentByTestId(ctx context.Context, testId int) (*entities.TestContent, error) {
	version, ok := ctx.Value("version").(string)
	if !ok || version == "v1" {
		return r.FindContentFromDBByTestId(ctx, testId)
	}

	content, err := r.FindContentFromRedisByTestId(ctx, testId)
	if err == nil && content != nil {
		return content, nil
	}
	r.findContentByTestId.Lock()
	defer r.findContentByTestId.Unlock()
	content, err = r.FindContentFromDBByTestId(ctx, testId)
	if err != nil {
		return nil, err
	}
	go r.SaveTestContentToRedis(ctx, content)
	return content, nil
}

func (r *repo) FindContentFromDBByTestId(ctx context.Context, testId int) (*entities.TestContent, error) {
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

func (r *repo) FindContentFromRedisByTestId(ctx context.Context, testId int) (*entities.TestContent, error) {
	status := r.redis.Get(ctx, fmt.Sprintf("test_content_%d", testId))
	if status.Err() != nil {
		return nil, status.Err()
	}
	var content entities.TestContent
	err := json.NewDecoder(strings.NewReader(status.Val())).Decode(&content)
	if err != nil {
		return nil, err
	}
	return &content, err
}

func (r *repo) SaveTestContentToRedis(ctx context.Context, content *entities.TestContent) error {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(content)
	if err != nil {
		return err
	}
	status := r.redis.Set(ctx, fmt.Sprintf("test_content_%d", content.TestId), b.String(), 1*time.Hour)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}
