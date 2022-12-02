package test_repository

import (
	"context"
	"picket/src/entities"
)

func (r *repo) FindTestByUserId(ctx context.Context, userId int) ([]entities.Test, error) {
	db := r.GetDB(ctx)

	list := make([]model, 0)
	if err := db.WithContext(ctx).Where("created_by = ?", userId).Find(&list).Error; err != nil {
		return nil, err
	}

	result := make([]entities.Test, len(list))

	for index, item := range list {
		result[index] = entities.Test{
			Id:                 item.Id,
			Code:               item.Code,
			Name:               item.Name,
			TimeToDo:           item.TimeToDo,
			TimeStart:          item.TimeStart,
			TimeEnd:            item.TimeEnd,
			DoOnce:             item.DoOnce,
			Password:           item.Password,
			PreventCheat:       item.PreventCheat,
			IsAuthenticateUser: item.IsAuthenticateUser,
			ShowMark:           item.ShowMark,
			ShowAnswer:         item.ShowAnswer,
			CreatedBy:          item.CreatedBy,
			CreatedAt:          item.CreatedAt,
			UpdatedAt:          item.UpdatedAt,
			DeletedAt:          &item.DeletedAt.Time,
			Version:            item.Version,
		}
	}

	return result, nil

}
