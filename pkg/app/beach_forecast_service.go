package app

import (
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app/timeforecast"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
)

type BeachForecastService struct {
	stormglassClient StormglassClient
}

func NewBeachForecastService(stormglassClient StormglassClient) *BeachForecastService {
	return &BeachForecastService{
		stormglassClient: stormglassClient,
	}
}

func (service *BeachForecastService) GetBeachForecasts(beaches []entity.Beach) ([]timeforecast.TimeForecast, error) {
	timeForecastsBuilder := timeforecast.NewTimeForecastsBuilder()
	for _, beach := range beaches {
		points, err := service.stormglassClient.FetchPoints(beach.Lat, beach.Lng)
		if err != nil {
			return nil, err
		}
		for _, point := range points {
			timeForecastsBuilder.BeachForecast(point.Time, timeforecast.NewBeachForecast(beach, point))
		}
	}
	return timeForecastsBuilder.Build(), nil
}
