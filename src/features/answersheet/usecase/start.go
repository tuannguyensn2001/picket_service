package answersheet_usecase

import (
	"context"
	"go.uber.org/zap"
	"picket/src/entities"
	answersheet_struct "picket/src/features/answersheet/struct"
	errpkg "picket/src/packages/err"
	"time"
)

func (u *usecase) Start(ctx context.Context, testId int, userId int) (*answersheet_struct.StartOutput, error) {
	checkDoing, err := u.CheckUserDoingTest(ctx, userId, testId)
	if err != nil {
		return nil, err
	}
	if checkDoing {
		return nil, errpkg.Answersheet.UserDoingTest
	}
	test, err := u.testUsecase.GetById(ctx, testId)
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
	err = u.repository.SendEvent(ctx, &event)
	if err != nil {
		return nil, err
	}

	content, err := u.testUsecase.GetContent(ctx, testId)
	if err != nil {
		return nil, err
	}

	zap.S().Info(content)

	return nil, nil
}
