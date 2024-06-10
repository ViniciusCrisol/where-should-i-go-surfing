package app

import (
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

func (service *RatingService) GetRatingBasedOnWindAndWavePositions(windPosition, wavePosition position.Position) int {
	if service.isWindOnshore(windPosition, wavePosition) {
		return 1
	}
	if service.isWindOffshore(windPosition, wavePosition) {
		return 5
	}
	return 3
}

func (service *RatingService) isWindOnshore(windPosition, wavePosition position.Position) bool {
	return windPosition == wavePosition
}

func (service *RatingService) isWindOffshore(windPosition, wavePosition position.Position) bool {
	return (service.beach.Position == position.N && wavePosition == position.N && windPosition == position.S) ||
		(service.beach.Position == position.S && wavePosition == position.S && windPosition == position.N) ||
		(service.beach.Position == position.E && wavePosition == position.E && windPosition == position.W) ||
		(service.beach.Position == position.W && wavePosition == position.W && windPosition == position.E)
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
