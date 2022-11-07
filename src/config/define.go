package config

import "gorm.io/gorm"

type IConfig interface {
	GetEnv() string
	GetGrpcAddress() string
	GetHttpAddress() string
	GetGoogleClientId() string
	GetGoogleClientSecret() string
	GetClientUrl() string
	GetDB() *gorm.DB
	GetSecretKey() string
}

func (c config) GetGoogleClientId() string {
	return c.GoogleClientId
}

func (c config) GetGoogleClientSecret() string {
	return c.GoogleClientSecret
}

func (c config) GetClientUrl() string {
	return c.ClientUrl
}

func (c config) GetSecretKey() string {
	return c.secretKey
}
