package app

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app/point"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app/timeforecast"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/test/mock/app/stormglassclient"
)

func TestBeachForecastService_GetBeachForecasts(t *testing.T) {
	beach1 := entity.Beach{
		Lat:      1.1,
		Lng:      1.1,
		Name:     "***",
		Position: entity.N,
	}
	beach2 := entity.Beach{
		Lat:      1.11,
		Lng:      1.11,
		Name:     "****",
		Position: entity.S,
	}
	point1 := point.Point{
		Time:           time.Date(2020, 04, 26, 00, 00, 00, 00, time.FixedZone("", 0)),
		SwellDirection: 64.26, SwellHeight: 0.15, SwellPeriod: 3.89,
		WaveDirection: 23.38, WaveHeight: 0.47,
		WindDirection: 29.45, WindSpeed: 10.1,
	}
	point2 := point.Point{
		Time:           time.Date(2020, 04, 26, 01, 00, 00, 00, time.FixedZone("", 0)),
		SwellDirection: 12.4, SwellHeight: 0.21, SwellPeriod: 3.67,
		WaveDirection: 23.1, WaveHeight: 0.46,
		WindDirection: 31.4, WindSpeed: 10.1,
	}
	point3 := point.Point{
		Time:           time.Date(2020, 04, 26, 00, 00, 00, 00, time.FixedZone("", 0)),
		SwellDirection: 94.21, SwellHeight: 0.12, SwellPeriod: 3.26,
		WaveDirection: 11.21, WaveHeight: 1.21,
		WindDirection: 21.33, WindSpeed: 2.43,
	}
	point4 := point.Point{
		Time:           time.Date(2020, 04, 26, 01, 00, 00, 00, time.FixedZone("", 0)),
		SwellDirection: 34.4, SwellHeight: 0.45, SwellPeriod: 3.66,
		WaveDirection: 45.11, WaveHeight: 3.46,
		WindDirection: 31.4, WindSpeed: 10.22,
	}

	var (
		mockedStormglassClient *stormglassclient.StormglassClient
		beachForecastService   *BeachForecastService
	)

	setup := func() {
		mockedStormglassClient = &stormglassclient.StormglassClient{}
		beachForecastService = NewBeachForecastService(mockedStormglassClient)
	}

	t.Run(
		"It should return an empty slice when no beaches are provided", func(t *testing.T) {
			setup()

			timeForecasts, err := beachForecastService.GetBeachForecasts(nil)

			assert.NoError(t, err)
			assert.Empty(t, timeForecasts)
		},
	)

	t.Run(
		"It should return multiple forecasts with different times for the same beach", func(t *testing.T) {
			setup()
			mockedStormglassClient.On("FetchPoints", beach1.Lat, beach1.Lng).Return([]point.Point{point1, point2}, nil)

			timeForecasts, err := beachForecastService.GetBeachForecasts([]entity.Beach{beach1})

			assert.NoError(t, err)
			assert.Len(t, timeForecasts, 2)
			assert.Len(t, timeForecasts[0].Forecasts, 1)
			assert.Len(t, timeForecasts[1].Forecasts, 1)
			assert.Equal(t, point1.Time, timeForecasts[0].Time)
			assert.Equal(t, point2.Time, timeForecasts[1].Time)
			assert.Equal(t, timeforecast.NewBeachForecast(beach1, point1), timeForecasts[0].Forecasts[0])
			assert.Equal(t, timeforecast.NewBeachForecast(beach1, point2), timeForecasts[1].Forecasts[0])
		},
	)

	t.Run(
		"It should return multiple forecasts with different times for different beaches", func(t *testing.T) {
			setup()
			mockedStormglassClient.On("FetchPoints", beach1.Lat, beach1.Lng).Return([]point.Point{point1, point2}, nil)
			mockedStormglassClient.On("FetchPoints", beach2.Lat, beach2.Lng).Return([]point.Point{point3, point4}, nil)

			timeForecasts, err := beachForecastService.GetBeachForecasts([]entity.Beach{beach1, beach2})

			assert.NoError(t, err)
			assert.Len(t, timeForecasts, 2)
			assert.Len(t, timeForecasts[0].Forecasts, 2)
			assert.Len(t, timeForecasts[1].Forecasts, 2)
			assert.Equal(t, point1.Time, timeForecasts[0].Time)
			assert.Equal(t, point3.Time, timeForecasts[0].Time)
			assert.Equal(t, point2.Time, timeForecasts[1].Time)
			assert.Equal(t, point4.Time, timeForecasts[1].Time)
			assert.Equal(t, timeforecast.NewBeachForecast(beach1, point1), timeForecasts[0].Forecasts[0])
			assert.Equal(t, timeforecast.NewBeachForecast(beach2, point3), timeForecasts[0].Forecasts[1])
			assert.Equal(t, timeforecast.NewBeachForecast(beach1, point2), timeForecasts[1].Forecasts[0])
			assert.Equal(t, timeforecast.NewBeachForecast(beach2, point4), timeForecasts[1].Forecasts[1])
		},
	)

	t.Run(
		"It should return an error when stormglass client returns an error", func(t *testing.T) {
			setup()
			mockedStormglassClient.On("FetchPoints", beach1.Lat, beach1.Lng).Return(nil, errors.ErrUnsupported)

			timeForecasts, err := beachForecastService.GetBeachForecasts([]entity.Beach{beach1})

			assert.Error(t, err)
			assert.Empty(t, timeForecasts)
		},
	)
}