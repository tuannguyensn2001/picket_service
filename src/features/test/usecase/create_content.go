package test_usecase

import (
	"context"
	"go.uber.org/zap"
	"picket/src/entities"
	test_struct "picket/src/features/test/struct"
)

func (u *usecase) CreateContent(ctx context.Context, input test_struct.CreateTestContentInput) error {
	var handler func(ctx context.Context, input test_struct.CreateTestContentInput) error
	switch input.Typeable {
	case 1:
		handler = u.CreateMultipleChoiceContent
	}

	err := handler(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (u *usecase) CreateMultipleChoiceContent(ctx context.Context, input test_struct.CreateTestContentInput) error {

	test, err := u.repository.FindByTestId(ctx, input.TestId)
	if err != nil {
		zap.S().Error(err)
		return err
	}

	//testContent, err := u.repository.FindContentByTestId(ctx, test.Id)
	//if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	//	zap.S().Error(err)
	//	return err
	//}
	//if testContent != nil {
	//	zap.S().Error(err)
	//	return errpkg.Test.TestHasContent
	//}

	ctx = u.repository.BeginTransaction(ctx)

	multipleChoice := entities.TestMultipleChoice{
		FilePath: input.MultipleChoice.FilePath,
		Score:    input.MultipleChoice.Score,
	}
	err = u.repository.CreateTestMultipleChoice(ctx, &multipleChoice)
	if err != nil {
		u.repository.Rollback(ctx)
		zap.S().Error(err)
		return err
	}

	answers := make([]entities.TestMultipleChoiceAnswer, len(input.MultipleChoice.Answers))
	for index, item := range input.MultipleChoice.Answers {
		answers[index] = entities.TestMultipleChoiceAnswer{
			Answer:               item.Answer,
			Score:                item.Score,
			Type:                 int(item.Type),
			TestMultipleChoiceId: multipleChoice.Id,
		}
	}
	err = u.repository.CreateListTestMultipleChoiceAnswers(ctx, answers)
	if err != nil {
		u.repository.Rollback(ctx)
		zap.S().Error(err)
		return err
	}

	content := entities.TestContent{
		TypeableId: multipleChoice.Id,
		Typeable:   1,
		TestId:     test.Id,
	}
	err = u.repository.CreateTestContent(ctx, &content)
	if err != nil {
		u.repository.Rollback(ctx)
		zap.S().Error(err)
		return err
	}

	//u.repository.Commit(ctx)

	return nil
}
