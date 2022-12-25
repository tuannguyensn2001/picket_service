package job_transport

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type IUsecase interface {
	UpdateSuccess(ctx context.Context, jobId int) error
	UpdateFail(ctx context.Context, jobId int, errFail error) error
}

type transport struct {
	usecase IUsecase
}

func New(ctx context.Context, usecase IUsecase) *transport {
	t := transport{usecase: usecase}
	go t.JobSuccess(ctx)
	go t.JobFail(ctx)
	return &t
}

func (t *transport) JobSuccess(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "job-success",
		GroupID: "consumer-picket-job-transport",
	})

	for {
		m, err := r.ReadMessage(ctx)
		zap.S().Info(m)
		if err != nil {
			zap.S().Error(err)
			continue
		}

		type SuccessPayload struct {
			JobId int `json:"job_id"`
		}
		var input SuccessPayload
		err = json.NewDecoder(bytes.NewBuffer(m.Value)).Decode(&input)
		if err != nil {
			zap.S().Error(err)
			continue
		}
		t.usecase.UpdateSuccess(ctx, input.JobId)
	}
}

func (t *transport) JobFail(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "job-fail",
		GroupID: "consumer-picket-job-transport-fail",
	})

	for {
		m, err := r.ReadMessage(ctx)
		zap.S().Info(m)
		if err != nil {
			zap.S().Error(err)
			continue
		}

		type FailPayload struct {
			JobId        int    `json:"job_id"`
			ErrorMessage string `json:"error_message"`
		}
		var input FailPayload
		err = json.NewDecoder(bytes.NewBuffer(m.Value)).Decode(&input)
		if err != nil || input.JobId == 0 {
			zap.S().Error(err)
			continue
		}
		t.usecase.UpdateFail(ctx, input.JobId, errors.New(input.ErrorMessage))
	}
}
