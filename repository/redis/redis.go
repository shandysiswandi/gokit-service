package redis

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	client *redis.Client
	mu     sync.RWMutex
}

type Configuration struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func NewRedis(conf Configuration) (*redisCache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.DB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &redisCache{client: rdb}, nil
}

func (r *redisCache) Close() error {
	return r.client.Close()
}
