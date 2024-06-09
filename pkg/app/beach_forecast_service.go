package app

import (
	"golang.org/x/sync/errgroup"

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
	var errGroup errgroup.Group
	timeForecastsBuilder := timeforecast.NewTimeForecastsBuilder()
	for _, beach := range beaches {
		errGroup.Go(
			func() error {
				points, err := service.stormglassClient.FetchPoints(beach.Lat, beach.Lng)
				if err != nil {
					return err
				}
				for _, point := range points {
					timeForecastsBuilder.BeachForecast(point.Time, timeforecast.NewBeachForecast(beach, point))
				}
				return nil
			},
		)
	}
	if err := errGroup.Wait(); err != nil {
		return nil, err
	}
	return timeForecastsBuilder.Build(), nil
}
