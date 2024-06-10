package timeforecast

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity/position"
)

func TestTimeForecastsBuilder(t *testing.T) {
	t.Run(
		"It should return an empty slice when no forecasts are added", func(t *testing.T) {
			assert.Empty(t, NewTimeForecastsBuilder().Build())
		},
	)

	t.Run(
		"It should add multiple forecasts with the same time", func(t *testing.T) {
			expectedForecast1 := BeachForecast{
				Time:           time.Date(2020, 04, 26, 00, 00, 00, 00, time.FixedZone("", 0)),
				Lat:            1.1,
				Lng:            1.1,
				Name:           "***",
				Position:       position.N.String(),
				SwellDirection: 64.26, SwellHeight: 0.15, SwellPeriod: 3.89,
				WaveDirection: 23.38, WaveHeight: 0.47,
				WindDirection: 29.45, WindSpeed: 10.1,
			}
			expectedForecast2 := BeachForecast{
				Time:           time.Date(2020, 04, 26, 00, 00, 00, 00, time.FixedZone("", 0)),
				Lat:            1.11,
				Lng:            1.11,
				Name:           "****",
				Position:       position.S.String(),
				SwellDirection: 12.4, SwellHeight: 0.21, SwellPeriod: 3.67,
				WaveDirection: 23.1, WaveHeight: 0.46,
				WindDirection: 31.4, WindSpeed: 10.1,
			}

			timeForecast := NewTimeForecastsBuilder().
				BeachForecast(expectedForecast1.Time, expectedForecast1).
				BeachForecast(expectedForecast2.Time, expectedForecast2).
				Build()

			assert.Len(t, timeForecast, 1)
			assert.Len(t, timeForecast[0].Forecasts, 2)
			assert.Equal(t, expectedForecast1.Time, timeForecast[0].Time)
			assert.Equal(t, expectedForecast1, timeForecast[0].Forecasts[0])
			assert.Equal(t, expectedForecast2, timeForecast[0].Forecasts[1])
		},
	)

	t.Run(
		"It should add multiple forecasts with different times", func(t *testing.T) {
			expectedForecast1 := BeachForecast{
				Time:           time.Date(2020, 04, 26, 00, 00, 00, 00, time.FixedZone("", 0)),
				Lat:            1.1,
				Lng:            1.1,
				Name:           "***",
				Position:       position.N.String(),
				SwellDirection: 64.26, SwellHeight: 0.15, SwellPeriod: 3.89,
				WaveDirection: 23.38, WaveHeight: 0.47,
				WindDirection: 29.45, WindSpeed: 10.1,
			}
			expectedForecast2 := BeachForecast{
				Time:           time.Date(2020, 04, 26, 01, 00, 00, 00, time.FixedZone("", 0)),
				Lat:            1.11,
				Lng:            1.11,
				Name:           "****",
				Position:       position.S.String(),
				SwellDirection: 12.4, SwellHeight: 0.21, SwellPeriod: 3.67,
				WaveDirection: 23.1, WaveHeight: 0.46,
				WindDirection: 31.4, WindSpeed: 10.1,
			}

			timeForecast := NewTimeForecastsBuilder().
				BeachForecast(expectedForecast1.Time, expectedForecast1).
				BeachForecast(expectedForecast2.Time, expectedForecast2).
				Build()

			assert.Len(t, timeForecast, 2)
			assert.Len(t, timeForecast[0].Forecasts, 1)
			assert.Len(t, timeForecast[1].Forecasts, 1)
			assert.Equal(t, expectedForecast1.Time, timeForecast[0].Time)
			assert.Equal(t, expectedForecast2.Time, timeForecast[1].Time)
			assert.Equal(t, expectedForecast1, timeForecast[0].Forecasts[0])
			assert.Equal(t, expectedForecast2, timeForecast[1].Forecasts[0])
		},
	)
}
