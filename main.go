package main

import (
	"github.com/my-stocks-pro/redis-service/infrastructure"
	"github.com/my-stocks-pro/redis-service/engine"
	"fmt"
	"os"
	"os/signal"
	"time"
	"context"
	"net/http"
)

func main() {

	config := infrastructure.NewConfig()

	logger, err := infrastructure.NewLogger()
	if err != nil {
		fmt.Println(err)
		return
	}

	redis := infrastructure.NewRedis(config)

	service := engine.New(config, logger, redis)

	service.InitMux()

	serverHTTP := &http.Server{
		Addr:    config.Port,
		Handler: service.Engine,
	}

	go func() {
		if err := serverHTTP.ListenAndServe(); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}()

	signal.Notify(service.QuitOS, os.Interrupt)
	select {
	case <-service.QuitOS:
		logger.Info("Shutdown Server by OS signal...")
	case <-service.QuitRPC:
		logger.Info("Shutdown Server by RPC signal...")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := serverHTTP.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("Server Shutdown: %s", err.Error()))
	}

	logger.Info("Server exiting")
}
