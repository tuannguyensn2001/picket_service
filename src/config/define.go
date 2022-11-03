package config

type IConfig interface {
	GetEnv() string
	GetGrpcAddress() string
}
