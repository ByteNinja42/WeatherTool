package repository

import (
	"encoding/json"
	"fmt"

	"github.com/ByteNinja42/WeatherTool/models"
	"github.com/go-redis/redis"
)

type WeatherRepo struct {
	*redis.Client
}

func NewWeatherRepo(client *redis.Client) (WeatherRepo, error) {
	return WeatherRepo{client}, nil
}
func (wr WeatherRepo) GetCachedWeatherForecast(city string) (models.WeatherForecastRepo, error) {
	var weather models.WeatherForecastRepo
	forecast, err := wr.Get(city).Bytes()
	if err != nil {
		if err == redis.Nil {
			return models.WeatherForecastRepo{}, models.ErrForecastNotFound
		}
		return models.WeatherForecastRepo{}, fmt.Errorf("error getting value redis : %w", err)
	}
	err = json.Unmarshal(forecast, &weather)
	if err != nil {
		return models.WeatherForecastRepo{}, fmt.Errorf("error unmarshalling value : %w", err)
	}

	return weather, nil
}

func (wr WeatherRepo) StoreWeatherForecast(city string, forecast models.WeatherForecastRepo) error {
	jsonWeather, err := json.Marshal(forecast)
	if err != nil {
		return err
	}

	return wr.Set(city, jsonWeather, 0).Err()
}
