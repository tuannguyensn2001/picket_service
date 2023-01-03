package answersheet_usecase

import (
	"context"
	"picket/src/entities"
	errpkg "picket/src/packages/err"
)

func (u *usecase) GetContent(ctx context.Context, testId int, userId int) (*entities.TestContent, error) {
	check, err := u.CheckUserDoingTest(ctx, userId, testId)
	if err != nil {
		return nil, err
	}
	if !check {
		return nil, errpkg.General.Forbidden
	}
	content, err := u.testUsecase.GetContent(ctx, testId)
	if err != nil {
		return nil, err
	}

	return content, nil
}
