package answersheet_usecase

import (
	"context"
	answersheet_struct "picket/src/features/answersheet/struct"
	errpkg "picket/src/packages/err"
)

func (u *usecase) GetContent(ctx context.Context, testId int, userId int) (*answersheet_struct.GetContentOutput, error) {
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

	timeLeft, err := u.GetTimeLeft(ctx, testId, userId)
	if err != nil {
		return nil, err
	}

	output := answersheet_struct.GetContentOutput{
		Content:  content,
		TimeLeft: timeLeft,
	}
	return &output, nil
}
