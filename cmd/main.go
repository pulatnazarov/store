package main

import (
	"context"
	"test/api"
	"test/config"
	"test/pkg/logger"
	"test/service"
	"test/storage/postgres"
	"test/storage/redis"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.ServiceName)

	pgStore, err := postgres.New(context.Background(), cfg, log)
	if err != nil {
		log.Error("error while connecting to db", logger.Error(err))
		return
	}
	defer pgStore.Close()

	newRedis := redis.New(cfg)

	services := service.New(pgStore, log, newRedis)

	server := api.New(services, log)

	log.Info("Service is running on", logger.Int("port", 8080))
	if err = server.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
