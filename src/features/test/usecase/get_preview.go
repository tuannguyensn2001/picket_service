package test_usecase

import (
	"context"
	"picket/src/entities"
)

func (u *usecase) GetPreview(ctx context.Context, code string) (*entities.Test, error) {
	return u.repository.FindByCode(ctx, code)
}
