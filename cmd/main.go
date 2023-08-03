package main

import (
	"fmt"
	"log"

	"github.com/ByteNinja42/WeatherTool/config"
	"github.com/ByteNinja42/WeatherTool/internal/repository"
	"github.com/ByteNinja42/WeatherTool/internal/service"
)

func main() {
	cfg := config.NewRedisConfig()
	client, err := repository.RedisClientInit(cfg)
	if err != nil {
		log.Fatal(fmt.Errorf("err creating redis client: %w", err))
	}

	rep, err := repository.NewWeatherRepo(client)
	if err != nil {
		log.Fatal(err)
	}
	service := service.NewWeatherService(rep)
	forecast, err := service.GetCurrentWeatherForecast("Minsk")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(forecast)
}
