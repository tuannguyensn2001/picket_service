package answersheet_usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"picket/src/entities"
	answersheet_struct "picket/src/features/answersheet/struct"
	errpkg "picket/src/packages/err"
	"time"
)

func (u *usecase) UserAnswer(ctx context.Context, userId int, input answersheet_struct.UserAnswerInput) error {
	ctx, span := tracer.Start(ctx, "user answer")
	defer span.End()
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return err
	}
	check, err := u.CheckUserDoingTest(ctx, userId, input.TestId)
	if err != nil {
		return err
	}
	if !check {
		return errpkg.Answersheet.UserDoingTest
	}
	err = u.testUsecase.CheckTestCanDo(ctx, input.TestId)
	if err != nil {
		zap.S().Error(err)
		return err
	}
	err = u.testUsecase.CheckTestAndQuestionValid(ctx, input.TestId, input.QuestionId)
	if err != nil {
		zap.S().Error(err)
		return err
	}

	now := time.Now()
	event := entities.AnswerSheetEvent{
		UserId:         userId,
		TestId:         input.TestId,
		Event:          entities.ANSWER,
		QuestionId:     input.QuestionId,
		PreviousAnswer: input.PreviousAnswer,
		Answer:         input.Answer,
		CreatedAt:      &now,
		UpdatedAt:      &now,
	}
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(event); err != nil {
		return err
	}
	job := entities.Job{
		Payload: b.String(),
		Status:  entities.INIT,
		Topic:   "answer-test",
	}
	ctx = u.repository.BeginTransaction(ctx)
	ctx, span = tracer.Start(ctx, "create job")
	if err := u.jobUsecase.Create(ctx, &job); err != nil {
		span.End()
		u.repository.Rollback(ctx)
		return err
	}
	span.End()
	b.Reset()
	payload := PayloadAnswerTest{
		JobId:   job.Id,
		Payload: event,
	}
	if err := json.NewEncoder(b).Encode(payload); err != nil {
		u.repository.Rollback(ctx)
		return err
	}

	w := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "answer-test",
		AllowAutoTopicCreation: true,
		BatchSize:              1,
	}

	key := fmt.Sprintf("%d-%d", userId, input.TestId)
	ctx, span = tracer.Start(ctx, "push to kafka")
	if err := w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(key),
		Value: b.Bytes(),
	}); err != nil {
		span.End()
		u.repository.Rollback(ctx)
		return err
	}

	span.End()
	u.repository.Commit(ctx)

	return nil
}
