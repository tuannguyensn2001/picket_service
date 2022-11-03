package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type structure struct {
	App struct {
		Env         string `mapstructure:"env"`
		GrpcAddress string `mapstructure:"grpcAddress"`
		HttpAddress string `mapstructure:"httpAddress"`
	} `mapstructure:"app"`
}

func bootstrap() structure {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	viper.ReadInConfig()

	var structure structure
	err = viper.Unmarshal(&structure)
	if err != nil {
		log.Fatalln(err)
	}
	return structure
}
