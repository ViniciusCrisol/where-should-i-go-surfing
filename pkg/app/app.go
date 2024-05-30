package app

import "time"

type Point struct {
	Time           time.Time
	SwellDirection float64
	SwellHeight    float64
	SwellPeriod    float64
	WaveDirection  float64
	WaveHeight     float64
	WindDirection  float64
	WindSpeed      float64
}

type StormglassClient interface {
	FetchPoints(lat, lng float64) ([]Point, error)
}
