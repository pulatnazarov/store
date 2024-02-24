package service

import (
	"context"
	"test/pkg/logger"
	"test/storage"
	"time"
)

type redisService struct {
	storage storage.IStorage
	log     logger.ILogger
	redis   storage.IRedisStorage
}

func NewRedisService(storage storage.IStorage, log logger.ILogger, redis storage.IRedisStorage) redisService {
	return redisService{
		storage: storage,
		log:     log,
		redis:   redis,
	}
}

func (r redisService) SetX(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	return r.redis.SetX(ctx, key, value, duration)
}

func (r redisService) Get(ctx context.Context, key string) interface{} {
	return r.redis.Get(ctx, key)
}
