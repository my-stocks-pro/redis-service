package main

import (
	"github.com/my-stocks-pro/redis-service/infrastructure"
	"github.com/my-stocks-pro/redis-service/engine"
	"fmt"
	"os"
)

func main() {

	config := infrastructure.NewConfig()

	logger, err := infrastructure.NewLogger()
	if err != nil {
		fmt.Println(err)
		return
	}

	redis := infrastructure.NewRedis(config)

	server := engine.New(config, logger, redis)

	server.InitMux()

	if err := server.Engine.Run(config.Port); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
