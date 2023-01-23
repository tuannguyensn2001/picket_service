package answersheet_usecase

import (
	"context"
	"time"
)

func (u *usecase) GetTimeLeft(ctx context.Context, testId int, userId int) (*time.Duration, error) {

	test, err := u.testUsecase.GetById(ctx, testId)
	if err != nil {
		return nil, err
	}

	if test.TimeEnd != nil {
		left := test.TimeEnd.Sub(time.Now())
		return &left, nil
	}

	latest, err := u.GetLatestStartTime(ctx, testId, userId)
	if err != nil {
		return nil, err
	}
	left := latest.Add(time.Duration(test.TimeToDo) * time.Minute).Sub(time.Now())

	return &left, nil
}
