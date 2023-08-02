package main

import (
	"fmt"
	"log"

	"github.com/ByteNinja42/WeatherTool/internal/repository"
	"github.com/ByteNinja42/WeatherTool/internal/service"
)

func main() {
	rep := repository.NewWeatherRepo()
	service := service.NewWeatherService(rep)
	forecast, err := service.GetCurrentWeatherForecast("Boston")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(forecast)
}
