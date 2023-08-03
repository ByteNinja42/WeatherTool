package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ByteNinja42/WeatherTool/internal/entities"
)

type WeatherService struct {
	WeatherRepo
}

func NewWeatherService(weatherRepo WeatherRepo) WeatherService {
	return WeatherService{WeatherRepo: weatherRepo}
}

type WeatherRepo interface {
	GetCachedWeatherForecast(city string) (entities.WeatherForecastRepo, error)
	StoreWeatherForecast(city string, forecast entities.WeatherForecastRepo) error
}

func (w WeatherService) GetCurrentWeatherForecast(city string) (entities.WeatherForecast, error) {
	var forecastTime time.Time
	if city == "" {
		return entities.WeatherForecast{}, fmt.Errorf("city name can't be empty")
	}

	weatherForecastFull, err := w.GetCachedWeatherForecast(city)
	if err != nil {
		if !errors.Is(err, entities.ErrForecastNotFound) {
			return entities.WeatherForecast{}, err
		}

		weatherForecastFull, err = getForecastAPI(city)
		if err != nil {
			return entities.WeatherForecast{}, err
		}
		err = w.StoreWeatherForecast(city, weatherForecastFull)
		if err != nil {
			return entities.WeatherForecast{}, err
		}
		return weatherForecastToResponse(weatherForecastFull), nil
	}

	timeNow, err := getTimeNowLocation(weatherForecastFull.Location.TzID)
	if err != nil {
		return entities.WeatherForecast{}, err
	}

	location, err := time.LoadLocation(weatherForecastFull.Location.TzID)
	if err != nil {
		return entities.WeatherForecast{}, err
	}

	forecastTime, err = time.ParseInLocation("2006-01-02 15:04", weatherForecastFull.Location.Localtime, location)

	if forecastTime.Add(time.Hour).Before(timeNow) {
		weatherForecastFull, err = getForecastAPI(city)
		if err != nil {
			return entities.WeatherForecast{}, err
		}
		err = w.StoreWeatherForecast(city, weatherForecastFull)
		if err != nil {
			return entities.WeatherForecast{}, err
		}
	}
	return weatherForecastToResponse(weatherForecastFull), nil

}

func getForecastAPI(city string) (entities.WeatherForecastRepo, error) {
	var forecastFull entities.WeatherForecastRepo
	client := http.DefaultClient
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=ca67825b0cd44493a6c90650230208&q=%s&aqi=no", city)
	resp, err := client.Get(url)
	if err != nil {
		return entities.WeatherForecastRepo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResponse ErrorResponseAPI
		err = json.NewDecoder(resp.Body).Decode(&errResponse)
		if err != nil {
			return entities.WeatherForecastRepo{}, err
		}
		return entities.WeatherForecastRepo{}, fmt.Errorf("error request API : status code : %d message : %s", resp.StatusCode, errResponse.Error.Message)
	}

	err = json.NewDecoder(resp.Body).Decode(&forecastFull)
	if err != nil {
		return entities.WeatherForecastRepo{}, err
	}
	return forecastFull, nil
}

func getTimeNowLocation(timeZone string) (time.Time, error) {
	location, err := time.LoadLocation(timeZone)
	if err != nil {
		return time.Time{}, err
	}
	currentTime := time.Now().In(location)
	return currentTime, nil
}

func weatherForecastToResponse(fullForecast entities.WeatherForecastRepo) entities.WeatherForecast {

	forecast := entities.WeatherForecast{
		City:                   fullForecast.Location.Name,
		Country:                fullForecast.Location.Country,
		Temperature:            fullForecast.Current.TempC,
		TemperatureMeasurement: "°C",
		LocalTime:              fullForecast.Location.Localtime,
		Condition:              fullForecast.Current.Condition.Text,
		WindSpeed:              fullForecast.Current.WindKph,
		WindMeasurement:        "kph",
		Humidity:               fullForecast.Current.Humidity,
	}
	if fullForecast.Location.Country == "United States of America" {
		forecast.TemperatureMeasurement = "°F"
		forecast.WindMeasurement = "mph"
	}
	return forecast
}
