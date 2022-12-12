package test_repository

import (
	"context"
	"picket/src/entities"
)

func (r *repo) FindByCode(ctx context.Context, code string) (*entities.Test, error) {
	db := r.GetDB(ctx).WithContext(ctx)
	var model model

	if err := db.Where("code = ?", code).First(&model).Error; err != nil {
		return nil, err
	}

	result := entities.Test{
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
		ShowMark:           model.ShowMark,
		ShowAnswer:         model.ShowAnswer,
		CreatedBy:          model.CreatedBy,
		CreatedAt:          model.CreatedAt,
		UpdatedAt:          model.UpdatedAt,
		DeletedAt:          &model.DeletedAt.Time,
		Version:            model.Version,
	}

	return &result, nil
}
