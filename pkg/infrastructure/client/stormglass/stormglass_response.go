package stormglass

import "time"

type FetchPointsResponse struct {
	Points []Point `json:"hours"`
}

type Point struct {
	Time           time.Time
	SwellDirection map[string]float64 `json:"swellDirection"`
	SwellHeight    map[string]float64 `json:"swellHeight"`
	SwellPeriod    map[string]float64 `json:"swellPeriod"`
	WaveDirection  map[string]float64 `json:"waveDirection"`
	WaveHeight     map[string]float64 `json:"waveHeight"`
	WindDirection  map[string]float64 `json:"windDirection"`
	WindSpeed      map[string]float64 `json:"windSpeed"`
}

func (point *Point) IsValid() bool {
	_, swellDirectionOK := point.SwellDirection["noaa"]
	_, swellHeightOK := point.SwellHeight["noaa"]
	_, swellPeriodOK := point.SwellPeriod["noaa"]
	_, waveDirectionOK := point.WaveDirection["noaa"]
	_, waveHeightOK := point.WaveHeight["noaa"]
	_, windDirectionOK := point.WindDirection["noaa"]
	_, windSpeedOK := point.WindSpeed["noaa"]
	return swellDirectionOK && swellHeightOK && swellPeriodOK &&
		waveDirectionOK && waveHeightOK &&
		windDirectionOK && windSpeedOK &&
		!point.Time.IsZero()
}
