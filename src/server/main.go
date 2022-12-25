package main

import (
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"log"
	"math/rand"
	"picket/src/cmd"
	"picket/src/config"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cfg := config.GetConfig()

	logger := getLogger(cfg)
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	otelLog := otelzap.New(logger)
	defer otelLog.Sync()
	undo := otelzap.ReplaceGlobals(otelLog)
	defer undo()

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
