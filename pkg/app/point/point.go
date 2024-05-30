package point

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
