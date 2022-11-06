package user_repository

import (
	"gorm.io/gorm"
	"time"
)

type user struct {
	Id        int        `gorm:"column:id"`
	Username  string     `gorm:"column:username"`
	Email     string     `gorm:"column:email"`
	Password  string     `gorm:"column:password"`
	Type      int        `gorm:"column:type"`
	Status    int        `gorm:"column:status"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt *gorm.DeletedAt
}

type profile struct {
	Id        int        `gorm:"column:id"`
	UserId    int        `gorm:"column:user_id"`
	Avatar    string     `gorm:"column:avatar"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt *gorm.DeletedAt
}

func (user) TableName() string {
	return "users"
}

func (profile) TableName() string {
	return "profiles"
}

const (
	type_account_google = 2
	type_account_normal = 1
	active              = 1
)
