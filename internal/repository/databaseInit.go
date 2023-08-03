package repository

import (
	"fmt"

	"github.com/ByteNinja42/WeatherTool/config"
	"github.com/go-redis/redis"
)

func RedisClientInit() (*redis.Client, error) {

	cfg := config.NewRedisConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	err := client.Ping().Err()
	if err != nil {
		return &redis.Client{}, fmt.Errorf("error ping redis %w", err)
	}
	return client, nil
}
