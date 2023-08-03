package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ByteNinja42/WeatherTool/config"
	"github.com/ByteNinja42/WeatherTool/internal/repository"
	"github.com/ByteNinja42/WeatherTool/internal/service"
)

func main() {
	cityToSearch := "Warsaw"
	if len(os.Args) > 1 {
		cityToSearch = os.Args[1]
	}
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

	forecast, err := service.GetCurrentWeatherForecast(cityToSearch)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", forecast)
}
