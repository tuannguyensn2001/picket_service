package test_usecase

import (
	"context"
	"picket/src/entities"
	errpkg "picket/src/packages/err"
)

func (u *usecase) CheckTestAndQuestionValid(ctx context.Context, testId int, questionId int) error {
	content, err := u.repository.FindContentByTestId(ctx, testId)
	if err != nil {
		return err
	}
	if content.Typeable == entities.MULTIPLE_CHOICE {
		questions, err := u.repository.FindTestMultipleChoiceAnswer(ctx, content.TypeableId)
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
