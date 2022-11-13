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
	Code        string `gorm:"column:code"`
}

type member struct {
	Id        int        `gorm:"column:id"`
	UserId    int        `gorm:"column:user_id"`
	ClassId   int        `gorm:"column:class_id"`
	Role      int        `gorm:"column:role"`
	Status    int        `gorm:"column:status"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (member) TableName() string {
	return "members"
}

func (model) TableName() string {
	return "classes"
}

const (
	teacher = 1
	student = 2
	active  = 1
	pending = 2
	removed = 3
)
