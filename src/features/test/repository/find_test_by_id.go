package test_repository

import (
	"context"
	"picket/src/entities"
)

func (r *repo) FindByTestId(ctx context.Context, id int) (*entities.Test, error) {
	var model model
	db := r.GetDB(ctx)
	if err := db.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	test := entities.Test{
		Id:                 model.Id,
		Code:               model.Code,
		Name:               model.Name,
		TimeToDo:           model.TimeToDo,
		TimeStart:          model.TimeStart,
		TimeEnd:            model.TimeEnd,
		DoOnce:             model.DoOnce,
		Password:           model.Password,
		PreventCheat:       model.PreventCheat,
		IsAuthenticateUser: model.IsAuthenticateUser,
		ShowAnswer:         model.ShowAnswer,
		ShowMark:           model.ShowMark,
		CreatedBy:          model.CreatedBy,
		CreatedAt:          model.CreatedAt,
		UpdatedAt:          model.UpdatedAt,
		Version:            model.Version,
	}

	return &test, nil
}
