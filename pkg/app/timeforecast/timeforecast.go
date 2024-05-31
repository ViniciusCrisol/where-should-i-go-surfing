package timeforecast

import (
	"time"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app/point"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
)

type TimeForecast struct {
	Time      time.Time       `json:"time"`
	Forecasts []BeachForecast `json:"forecasts"`
}

type BeachForecast struct {
	Time           time.Time `json:"time"`
	Lat            float64   `json:"lat"`
	Lng            float64   `json:"lng"`
	Name           string    `json:"name"`
	Position       string    `json:"position"`
	SwellDirection float64   `json:"swell_direction"`
	SwellHeight    float64   `json:"swell_height"`
	SwellPeriod    float64   `json:"swell_period"`
	WaveDirection  float64   `json:"wave_direction"`
	WaveHeight     float64   `json:"wave_height"`
	WindDirection  float64   `json:"wind_direction"`
	WindSpeed      float64   `json:"wind_speed"`
}

func NewBeachForecast(beach entity.Beach, point point.Point) BeachForecast {
	return BeachForecast{
		Time:           point.Time,
		Lat:            beach.Lat,
		Lng:            beach.Lng,
		Name:           beach.Name,
		Position:       beach.Position.String(),
		SwellDirection: point.SwellDirection,
		SwellHeight:    point.SwellHeight,
		SwellPeriod:    point.SwellPeriod,
		WaveDirection:  point.WaveDirection,
		WaveHeight:     point.WaveHeight,
		WindDirection:  point.WindDirection,
		WindSpeed:      point.WindSpeed,
	}
}
