package class_repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"picket/src/entities"
	class_struct "picket/src/features/class/struct"
	"picket/src/repository"
)

type repo struct {
	repository.Repository
}

func New(db *gorm.DB) *repo {
	return &repo{
		repository.Repository{
			Db: db,
		},
	}
}

func (r *repo) Create(ctx context.Context, class *entities.Class) error {
	db := r.GetDB(ctx)
	model := model{
		Name:        class.Name,
		Description: class.Description,
		Code:        class.Code,
	}
	if err := db.WithContext(ctx).Create(&model).Error; err != nil {
		return err
	}
	class.Id = model.Id
	class.CreatedAt = model.CreatedAt
	class.UpdatedAt = model.UpdatedAt
	return nil
}

func (r *repo) AddTeacherToClass(ctx context.Context, userId int, classId int) error {
	db := r.GetDB(ctx)
	model := member{
		UserId:  userId,
		ClassId: classId,
		Role:    teacher,
		Status:  active,
	}
	if err := db.WithContext(ctx).Create(&model).Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) GetAllCodes(ctx context.Context) ([]string, error) {
	db := r.GetDB(ctx)
	result := make([]string, 0)
	if err := db.WithContext(ctx).Raw("select code from classes").Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repo) FindByQuery(ctx context.Context, params class_struct.QueryClass) ([]entities.Class, error) {
	db := r.GetDB(ctx).WithContext(ctx)
	result := make([]model, 0)
	query := db.Model(&model{})
	if len(params.Name) > 0 {
		query = query.Where("name like %d?%d", params.Name)
	}
	if len(params.OrderBy) > 0 && len(params.Direction) > 0 {
		query = query.Order(fmt.Sprintf("%s %s", params.OrderBy, params.Direction))
	}
	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}

	resp := make([]entities.Class, len(result))

	for index, item := range result {
		resp[index] = entities.Class{
			Name:        item.Name,
			Id:          item.Id,
			Description: item.Description,
			Code:        item.Code,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}
	}

	return resp, nil
}
