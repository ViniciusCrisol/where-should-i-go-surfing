package app

import (
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app/dto/point"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity/position"
)

type RatingService struct {
	beach entity.Beach
}

func NewRatingService(beach entity.Beach) *RatingService {
	return &RatingService{
		beach: beach,
	}
}

func (service *RatingService) GetRating(point point.Point) int {
	windDirection := service.GetPositionFromLocation(point.WindDirection)
	waveDirection := service.GetPositionFromLocation(point.WaveDirection)
	windAndWaveRating := service.GetRatingBasedOnWindAndWaveDirections(windDirection, waveDirection)
	swellPeriodRating := service.GetRatingBasedOnSwellPeriod(point.SwellPeriod)
	swellHeightRating := service.GetRatingBasedOnSwellHeight(point.SwellHeight)
	return (windAndWaveRating + swellPeriodRating + swellHeightRating) / 3
}

func (service *RatingService) GetPositionFromLocation(location float64) position.Position {
	if (location >= 0 && location < 50) || location >= 310 {
		return position.N
	}
	if location >= 50 && location < 120 {
		return position.E
	}
	if location >= 120 && location < 220 {
		return position.S
	}
	return position.W
}

func (service *RatingService) GetRatingBasedOnWindAndWaveDirections(windDirection, waveDirection position.Position) int {
	if service.isWindOnshore(windDirection, waveDirection) {
		return 1
	}
	if service.isWindOffshore(windDirection, waveDirection) {
		return 5
	}
	return 3
}

func (service *RatingService) isWindOnshore(windDirection, waveDirection position.Position) bool {
	return windDirection == waveDirection
}

func (service *RatingService) isWindOffshore(windDirection, waveDirection position.Position) bool {
	return (service.beach.Position == position.N && waveDirection == position.N && windDirection == position.S) ||
		(service.beach.Position == position.S && waveDirection == position.S && windDirection == position.N) ||
		(service.beach.Position == position.E && waveDirection == position.E && windDirection == position.W) ||
		(service.beach.Position == position.W && waveDirection == position.W && windDirection == position.E)
}

func (service *RatingService) GetRatingBasedOnSwellPeriod(period float64) int {
	if period < 7 {
		return 1
	}
	if period >= 7 && period < 9 {
		return 2
	}
	if period >= 10 && period < 14 {
		return 4
	}
	return 5
}

func (service *RatingService) GetRatingBasedOnSwellHeight(height float64) int {
	if height < 0.3 {
		return 1
	}
	if height >= 0.3 && height < 1 {
		return 2
	}
	if height >= 1 && height < 2 {
		return 3
	}
	return 5
}
