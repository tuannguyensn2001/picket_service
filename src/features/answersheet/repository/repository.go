package answersheet_repository

import (
	"context"
	"gorm.io/gorm"
	"picket/src/entities"
	repository2 "picket/src/repository"
)

type repository struct {
	repository2.Repository
}

func New(db *gorm.DB) *repository {
	return &repository{
		repository2.Repository{Db: db},
	}
}

func (r *repository)GetLatestEvent(ctx context.Context,userId int, testId int, number int) ([]entities.AnswerSheetEvent,error)  {
	return nil,nil
}