package entities

import "errors"

var ErrForecastNotFound = errors.New("forecast for this city wasn't found")
