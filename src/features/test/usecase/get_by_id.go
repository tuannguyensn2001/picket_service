package test_usecase

import (
	"context"
	"go.opentelemetry.io/otel"
	"picket/src/entities"
)

var tracer = otel.Tracer("test_usecase")

func (u *usecase) GetById(ctx context.Context, id int) (*entities.Test, error) {
	return u.repository.FindByTestId(ctx, id)
}

func (u *usecase) GetContent(ctx context.Context, testId int) (*entities.TestContent, error) {
	ctx, span := tracer.Start(ctx, "get content by test id")
	content, err := u.repository.FindContentByTestId(ctx, testId)
	span.End()
	if err != nil {
		return nil, err
	}
	ctx, span = tracer.Start(ctx, "get multiple choice")
	multipleChoice, err := u.repository.FindTestMultipleChoiceByTestId(ctx, testId)
	span.End()
	if err != nil {
		return nil, err
	}
	ctx, span = tracer.Start(ctx, "get multiple choice answers")
	answers, err := u.repository.FindTestMultipleChoiceAnswer(ctx, multipleChoice.Id)
	span.End()
	if err != nil {
		return nil, err
	}
	multipleChoice.Answers = answers
	content.MultipleChoice = multipleChoice

	return content, nil
}
