package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ByteNinja42/WeatherTool/config"
	"github.com/ByteNinja42/WeatherTool/internal/handlers"
	"github.com/ByteNinja42/WeatherTool/internal/repository"
	"github.com/ByteNinja42/WeatherTool/internal/service"
	"github.com/labstack/echo/v4"
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
	handler := handlers.NewHandler(service)
	if err := startServer(handler); err != nil {
		log.Fatal(err)
	}
}

func startServer(handler handlers.Handler) error {
	e := echo.New()
	e.GET("/forecast/current/:city", handler.GetWeather)
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		return err
	}
	return nil
}
