package main

import (
	"go.uber.org/zap"
	"log"
	"myclass_service/src/cmd"
	"myclass_service/src/config"
)

func main() {
	cfg := config.GetConfig()

	logger := getLogger(cfg)
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	err := cmd.GetRoot(cfg).Execute()
	if err != nil {
		zap.S().Fatalln(err)
	}

}

func getLogger(config config.IConfig) *zap.Logger {
	result, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln(err)
	}
	if config.GetEnv() == "PRODUCTION" {
		result, err = zap.NewProduction()
		if err != nil {
			log.Fatalln(err)
		}
	}

	return result
}
