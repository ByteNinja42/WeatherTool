package handlers

import (
	"fmt"
	"net/http"

	"github.com/ByteNinja42/WeatherTool/internal/entities"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	WeatherService
}

func NewHandler(weatherService WeatherService) Handler {
	return Handler{WeatherService: weatherService}
}

type WeatherService interface {
	GetCurrentWeatherForecast(city string) (entities.WeatherForecast, error)
}

func (h Handler) GetWeather(ctx echo.Context) error {
	fmt.Println("handler works")
	fmt.Println(ctx.Param("city"))
	forecast, err := h.GetCurrentWeatherForecast(ctx.Param("city"))
	if err != nil {
		return err
	}
	fmt.Println(forecast)
	return ctx.JSON(http.StatusOK, forecast)
}
