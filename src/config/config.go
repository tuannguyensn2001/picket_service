package config

import (
	"context"
	"github.com/go-redis/redis/v9"
	"gorm.io/driver/postgres"
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
	redis              *redis.Client
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

	db, err := gorm.Open(postgres.Open(structure.Database.Postgres), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}
	rd := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	status := rd.Ping(context.TODO())
	if status.Err() != nil {
		log.Println("redis ping error", status.Err())
	}

	result.db = db
	result.redis = rd

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
