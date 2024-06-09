package timeforecast

import (
	"sync"
	"time"
)

type TimeForecastsBuilder struct {
	mt            sync.Mutex
	timeForecasts []TimeForecast
}

func NewTimeForecastsBuilder() *TimeForecastsBuilder {
	return &TimeForecastsBuilder{
		timeForecasts: []TimeForecast{},
	}
}

func (builder *TimeForecastsBuilder) Build() []TimeForecast {
	return builder.timeForecasts
}

func (builder *TimeForecastsBuilder) BeachForecast(time time.Time, beachForecast BeachForecast) *TimeForecastsBuilder {
	builder.mt.Lock()
	defer builder.mt.Unlock()

	for i, timeForecast := range builder.timeForecasts {
		if timeForecast.Time.Equal(time) {
			builder.timeForecasts[i].Forecasts = append(builder.timeForecasts[i].Forecasts, beachForecast)
			return builder
		}
	}
	builder.timeForecasts = append(builder.timeForecasts, TimeForecast{time, []BeachForecast{beachForecast}})
	return builder
}
