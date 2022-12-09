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

func (r *repo) FindByTestId(ctx context.Context, id int) (*entities.Test, error) {

	status := r.redis.Get(ctx, fmt.Sprintf("test_%d", id))
	if status.Err() == nil {
		var test entities.Test
		r := strings.NewReader(status.Val())
		err := json.NewDecoder(r).Decode(&test)
		if err == nil {
			return &test, nil
		}
	}

	var model model
	db := r.GetDB(ctx)
	if err := db.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	test := entities.Test{
		Id:                 model.Id,
		Code:               model.Code,
		Name:               model.Name,
		TimeToDo:           model.TimeToDo,
		TimeStart:          model.TimeStart,
		TimeEnd:            model.TimeEnd,
		DoOnce:             model.DoOnce,
		Password:           model.Password,
		PreventCheat:       model.PreventCheat,
		IsAuthenticateUser: model.IsAuthenticateUser,
		ShowAnswer:         model.ShowAnswer,
		ShowMark:           model.ShowMark,
		CreatedBy:          model.CreatedBy,
		CreatedAt:          model.CreatedAt,
		UpdatedAt:          model.UpdatedAt,
		Version:            model.Version,
	}
	go func() {
		b := new(bytes.Buffer)
		err := json.NewEncoder(b).Encode(test)
		if err != nil {
			zap.S().Error(err)
		}
		status := r.redis.Set(ctx, fmt.Sprintf("test_%d", id), b.String(), 1*time.Hour)
		zap.S().Error(status.Err())
	}()

	return &test, nil
}
