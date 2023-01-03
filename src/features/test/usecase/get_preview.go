package test_usecase

import (
	"context"
	"picket/src/entities"
)

func (u *usecase) GetPreview(ctx context.Context, id int) (*entities.Test, error) {
	return u.repository.FindByTestId(ctx, id)
}
