package class_repository

import (
	"gorm.io/gorm"
	"time"
)

type model struct {
	Id          int        `gorm:"column:id"`
	Name        string     `gorm:"column:name"`
	Description string     `gorm:"column:description"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
	DeletedAt   *gorm.DeletedAt
}

func (model) TableName() string {
	return "classes"
}
