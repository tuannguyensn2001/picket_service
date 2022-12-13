package answersheet_usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"picket/src/entities"
	answersheet_struct "picket/src/features/answersheet/struct"
	errpkg "picket/src/packages/err"
	"time"
)

var tracer = otel.Tracer("answersheet_usecase")

func (u *usecase) Start(ctx context.Context, testId int, userId int) (*answersheet_struct.StartOutput, error) {
	ctx, span := tracer.Start(ctx, "start doing test")
	defer span.End()
	//checkDoing, err := u.CheckUserDoingTest(ctx, userId, testId)
	//if err != nil {
	//	return nil, err
	//}
	//if checkDoing {
	//	return nil, errpkg.Answersheet.UserDoingTest
	//}
	ctx, span = tracer.Start(ctx, "get test by id")
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

	event := entities.AnswerSheetEvent{
		UserId: userId,
		TestId: testId,
		Event:  entities.START,
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
