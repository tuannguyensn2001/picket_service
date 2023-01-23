package test_usecase

import (
	"context"
	"picket/src/entities"
	errpkg "picket/src/packages/err"
)

func (u *usecase) CheckTestAndQuestionValid(ctx context.Context, testId int, questionId int) error {
	ctx, span := tracer.Start(ctx, "get content by test id")
	defer span.End()
	content, err := u.repository.FindContentByTestId(ctx, testId)
	if err != nil {
		return err
	}
	if content.Typeable == entities.MULTIPLE_CHOICE {
		ctx, span = tracer.Start(ctx, "get test multiple choice answer")
		questions, err := u.repository.FindTestMultipleChoiceAnswer(ctx, content.TypeableId)
		span.End()
		if err != nil {
			return err
		}
		valid := false
		for _, item := range questions {
			if item.Id == questionId {
				valid = true
				break
			}
		}
		if !valid {
			return errpkg.Test.QuestionNotValid
		}
	}

	return nil
}
