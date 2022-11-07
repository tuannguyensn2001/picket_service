package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type config struct {
	Env                string
	GrpcAddress        string
	HttpAddress        string
	GoogleClientId     string
	GoogleClientSecret string
	ClientUrl          string
	db                 *gorm.DB
	secretKey          string
}

func GetConfig() config {
	structure := bootstrap()
	result := config{
		Env:                structure.App.Env,
		GrpcAddress:        structure.App.GrpcAddress,
		HttpAddress:        structure.App.HttpAddress,
		GoogleClientId:     structure.OAuth2.Google.ClientId,
		GoogleClientSecret: structure.OAuth2.Google.ClientSecret,
		ClientUrl:          structure.Client.Url,
		secretKey:          structure.App.SecretKey,
	}

	db, err := gorm.Open(mysql.Open(structure.Database.Mysql), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}

	result.db = db

	return result
}

func (c config) GetEnv() string {
	return c.Env
}

func (c config) GetGrpcAddress() string {
	return c.GrpcAddress
}

func (c config) GetHttpAddress() string {
	return c.HttpAddress
}

func (c config) GetDB() *gorm.DB {
	return c.db
}
