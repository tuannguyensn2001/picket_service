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
		SecretKey   string `mapstructure:"secretKey"`
	} `mapstructure:"app"`
	OAuth2 struct {
		Google struct {
			ClientId     string `mapstructure:"client_id"`
			ClientSecret string `mapstructure:"client_secret"`
		} `mapstructure:"google"`
	} `mapstructure:"oauth2"`
	Client struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"client"`
	Database struct {
		Mysql    string `mapstructure:"mysql"`
		Postgres string `mapstructure:"postgres"`
	} `mapstructure:"database"`
}

func bootstrap() structure {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetConfigName(".env.production")
	viper.ReadInConfig()

	viper.SetConfigName(".env")
	viper.MergeInConfig()

	for k, v := range bind {
		viper.Set(v, viper.GetString(k))
	}

	var structure structure
	err = viper.Unmarshal(&structure)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(structure)
	return structure
}
