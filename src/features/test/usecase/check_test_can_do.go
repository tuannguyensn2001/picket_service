package test_usecase

import (
	"context"
	errpkg "picket/src/packages/err"
	"time"
)

func (u *usecase) CheckTestCanDo(ctx context.Context, testId int) error {
	test, err := u.repository.FindByTestId(ctx, testId)
	if err != nil {
		return err
	}
	if test.TimeEnd != nil {
		if test.TimeEnd.Before(time.Now()) {
			return errpkg.Answersheet.TimeNotValid
		}
	}
	if test.TimeStart != nil {
		if test.TimeStart.After(time.Now()) {
			return errpkg.Answersheet.TimeNotValid
		}
	}

	return nil
}
