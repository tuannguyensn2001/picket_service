package user_repository

import (
	"context"
	"gorm.io/gorm"
	"myclass_service/src/entities"
	"myclass_service/src/repository"
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

func (r *repo) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	db := r.GetDB(ctx)
	var model user
	if err := db.WithContext(ctx).Where("email = ?", email).First(&model).Error; err != nil {
		return nil, err
	}

	result := entities.User{
		Email:    model.Email,
		Username: model.Username,
		Id:       model.Id,
		Status:   model.Status,
		Type:     model.Type,
		Password: model.Password,
	}
	return &result, nil
}

func (r *repo) FindById(ctx context.Context, id int) (*entities.User, error) {
	db := r.GetDB(ctx)

	type queryResult struct {
		Id       int    `gorm:"column:id"`
		Username string `gorm:"column:username"`
		Email    string `gorm:"column:email"`
		Avatar   string `gorm:"column:avatar"`
	}

	var result queryResult
	query := "SELECT users.id,users.username,users.email,profiles.avatar FROM users LEFT JOIN profiles ON  users.id = profiles.user_id WHERE users.id = ?"
	if err := db.WithContext(ctx).Raw(query, id).Scan(&result).Error; err != nil {
		return nil, err
	}

	user := entities.User{
		Email:    result.Email,
		Username: result.Username,
		Profile: &entities.Profile{
			Avatar: result.Avatar,
		},
	}

	return &user, nil
}

func (r *repo) CreateAccount(ctx context.Context, entity *entities.User) error {
	db := r.GetDB(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		model := user{
			Username: entity.Username,
			Password: entity.Password,
			Email:    entity.Email,
			Status:   active,
			Type:     type_account_normal,
		}
		if err := tx.WithContext(ctx).Create(&model).Error; err != nil {
			return err
		}

		profile := profile{
			UserId: model.Id,
			Avatar: default_avatar,
		}
		if err := tx.WithContext(ctx).Create(&profile).Error; err != nil {
			return err
		}
		entity.Status = model.Status
		entity.Type = model.Type
		return nil
	})

}

func (r *repo) InsertByGoogleAccount(ctx context.Context, entity *entities.User) error {
	db := r.GetDB(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		model := user{
			Username: entity.Username,
			Email:    entity.Email,
			Status:   active,
			Type:     type_account_google,
		}

		if err := tx.WithContext(ctx).Create(&model).Error; err != nil {
			return err
		}

		p := profile{
			UserId: model.Id,
			Avatar: entity.Profile.Avatar,
		}

		if err := tx.WithContext(ctx).Create(&p).Error; err != nil {
			return err
		}

		entity.Type = model.Type
		entity.Status = model.Status

		return nil
	})
}
