package test_repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"picket/src/entities"
	"strings"
	"time"
)

var tracer = otel.Tracer("test_repository")

func (r *repo) FindByTestId(ctx context.Context, id int) (*entities.Test, error) {
	version, ok := ctx.Value("version").(string)
	if !ok || version == "v1" {
		return r.FindTestFromDBById(ctx, id)
	}
	var test *entities.Test
	var err error
	
	test, err = r.FindTestFromRedisById(ctx, id)
	if test != nil && err == nil {
		return test, err
	}

	r.findTestById.Lock()
	defer r.findTestById.Unlock()
	test, err = r.FindTestFromDBById(ctx, id)
	if err != nil {
		return nil, err
	}
	go r.SaveTestToRedis(ctx, test)
	return test, nil
}

func (r *repo) FindTestFromRedisById(ctx context.Context, testId int) (*entities.Test, error) {
	ctx, span := tracer.Start(ctx, "find test from redis")
	defer span.End()
	status := r.redis.Get(ctx, fmt.Sprintf("test_%d", testId))
	if status.Err() == nil {
		var test entities.Test
		r := strings.NewReader(status.Val())
		err := json.NewDecoder(r).Decode(&test)
		if err == nil {
			return &test, nil
		}
	}
	return nil, status.Err()
}

func (r *repo) FindTestFromDBById(ctx context.Context, testId int) (*entities.Test, error) {
	ctx, span := tracer.Start(ctx, "find test from db")
	defer span.End()
	var model model
	db := r.GetDB(ctx)
	if errDb := db.WithContext(ctx).Where("id = ?", testId).First(&model).Error; errDb != nil {
		//return nil, err
		return nil, errDb
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
	return &test, nil
}

func (r *repo) SaveTestToRedis(ctx context.Context, test *entities.Test) error {
	ctx, span := tracer.Start(ctx, "save test to redis")
	defer span.End()
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(test)
	if err != nil {
		zap.S().Error(err)
	}
	var timeExpr time.Duration
	if test.TimeStart != nil && test.TimeEnd != nil {
		start := *test.TimeStart
		end := *test.TimeEnd
		timeExpr = start.Sub(end)
	} else {
		timeExpr = 1 * time.Hour
	}
	status := r.redis.Set(ctx, fmt.Sprintf("test_%d", test.Id), b.String(), timeExpr)
	return status.Err()
}
