package answersheet_usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"picket/src/entities"
	answersheet_struct "picket/src/features/answersheet/struct"
	errpkg "picket/src/packages/err"
	"time"
)

var tracer = otel.Tracer("answersheet_usecase")

type PayloadStartTest struct {
	JobId   int                       `json:"job_id"`
	Payload entities.AnswerSheetEvent `json:"payload"`
}

func (u *usecase) Start(ctx context.Context, testId int, userId int) (*answersheet_struct.StartOutput, error) {

	checkDoing, err := u.CheckUserDoingTest(ctx, userId, testId)
	if err != nil {
		return nil, err
	}
	if checkDoing {
		otelzap.Ctx(ctx).Error(errpkg.Answersheet.UserDoingTest.Message)
		return nil, errpkg.Answersheet.UserDoingTest
	}
	ctx, span := tracer.Start(ctx, "get test by id")
	test, err := u.testUsecase.GetById(ctx, testId)
	span.End()
	if err != nil {
		return nil, err
	}

	if test.TimeEnd != nil {
		if test.TimeEnd.Before(time.Now()) {
			return nil, errpkg.Answersheet.TimeNotValid
		}
	}
	if test.TimeStart != nil {
		if test.TimeStart.After(time.Now()) {
			return nil, errpkg.Answersheet.TimeNotValid
		}
	}

	now := time.Now()
	event := entities.AnswerSheetEvent{
		UserId:    userId,
		TestId:    testId,
		Event:     entities.START,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	//ctx, span = tracer.Start(ctx, "create event")
	//err = u.repository.SendEvent(ctx, &event)
	//span.End()
	//if err != nil {
	//	return nil, err
	//}

	ctx, span = tracer.Start(ctx, "get content")
	content, err := u.testUsecase.GetContent(ctx, testId)
	span.End()
	if err != nil {
		return nil, err
	}

	ctx, span = tracer.Start(ctx, "push to kafka")
	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(event)
	if err != nil {
		return nil, err
	}
	job := entities.Job{
		Payload: b.String(),
		Status:  entities.INIT,
	}
	err = u.jobUsecase.Create(ctx, &job)
	if err != nil {
		return nil, err
	}

	b = new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(PayloadStartTest{
		JobId:   job.Id,
		Payload: event,
	})
	if err != nil {
		return nil, err
	}

	w := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "start-test",
		AllowAutoTopicCreation: true,
		Balancer:               &kafka.LeastBytes{},
		BatchSize:              1,
	}
	err = w.WriteMessages(ctx, kafka.Message{
		Value: b.Bytes(),
	})
	span.End()
	if err != nil {
		return nil, err
	}
	zap.S().Info(content)

	return nil, nil
}
