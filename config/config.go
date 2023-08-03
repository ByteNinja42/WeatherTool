package config

import (
	"os"
	"strconv"
)

func getEnv(key, defaultValue string) string {
	env := os.Getenv(key)
	if env == "" {
		env = defaultValue
	}
	return env
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func NewRedisConfig() RedisConfig {
	dbNumber, err := strconv.Atoi(getEnv("REDISDB_NUMBER", "0"))
	if err != nil {
		dbNumber = 0
	}
	return RedisConfig{
		Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
		Password: getEnv("REDIS_PASSWORD", "qwerty"),
		DB:       dbNumber,
	}
}
