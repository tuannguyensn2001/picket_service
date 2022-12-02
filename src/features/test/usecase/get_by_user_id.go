package test_usecase

import (
	"context"
	"go.uber.org/zap"
	"picket/src/entities"
)

func (u *usecase) GetTestsByUserId(ctx context.Context, userId int) ([]entities.Test, error) {
	result, err := u.repository.FindTestByUserId(ctx, userId)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	return result, nil
}
