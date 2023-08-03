package repository

import "github.com/ByteNinja42/WeatherTool/models"

type WeatherRepo struct {
}

func NewWeatherRepo() WeatherRepo {
	return WeatherRepo{}
}
func (wr WeatherRepo) GetCachedWeatherForecast(city string) (models.WeatherForecastRepo, error) {
	return models.WeatherForecastRepo{}, models.ErrForecastNotFound
}
