package service

type WeatherForecast struct {
	City                   string  `json:"city"`
	Country                string  `json:"country"`
	Temperature            float64 `json:"temperature"`
	TemperatureMeasurement string  `json:"temperature_measurement"`
	LocalTime              string  `json:"local_time"`
	Condition              string  `json:"condition"`
	WindSpeed              float64 `json:"wind_kph"`
	WindMeasurement        string  `json:"wind_measurement"`
	Humidity               int64   `json:"humidity"`
}
