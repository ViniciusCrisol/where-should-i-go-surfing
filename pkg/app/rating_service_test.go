package app

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app/dto/point"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity/position"
)

var ratingService *RatingService

func setup() {
	ratingService = NewRatingService(
		entity.Beach{
			Lat:      1.1,
			Lng:      1.1,
			Name:     "***",
			Position: position.N,
		},
	)
}

func TestRatingService_GetRating(t *testing.T) {
	t.Run(
		"It should return 1 for a terrible point", func(t *testing.T) {
			setup()
			assert.Equal(
				t, 1, ratingService.GetRating(
					point.Point{
						SwellDirection: 64.26, SwellHeight: 0.15, SwellPeriod: 3.89,
						WaveDirection: 23.38, WaveHeight: 0.47,
						WindDirection: 19.45, WindSpeed: 10.1,
					},
				),
			)
		},
	)

	t.Run(
		"It should return 2 for a bad point", func(t *testing.T) {
			setup()
			assert.Equal(
				t, 2, ratingService.GetRating(
					point.Point{
						SwellDirection: 12.4, SwellHeight: 0.21, SwellPeriod: 3.67,
						WaveDirection: 23.1, WaveHeight: 0.46,
						WindDirection: 131.4, WindSpeed: 10.1,
					},
				),
			)
		},
	)

	t.Run(
		"It should return 3 for a regular point", func(t *testing.T) {
			setup()
			assert.Equal(
				t, 3, ratingService.GetRating(
					point.Point{
						SwellDirection: 12.4, SwellHeight: 0.31, SwellPeriod: 7.67,
						WaveDirection: 23.1, WaveHeight: 0.46,
						WindDirection: 131.4, WindSpeed: 10.1,
					},
				),
			)
		},
	)

	t.Run(
		"It should return 4 for a good point", func(t *testing.T) {
			setup()
			assert.Equal(
				t, 4, ratingService.GetRating(
					point.Point{
						SwellDirection: 12.4, SwellHeight: 0.31, SwellPeriod: 17.67,
						WaveDirection: 23.1, WaveHeight: 0.46,
						WindDirection: 131.4, WindSpeed: 10.1,
					},
				),
			)
		},
	)

	t.Run(
		"It should return 5 for an excellent point", func(t *testing.T) {
			setup()
			assert.Equal(
				t, 5, ratingService.GetRating(
					point.Point{
						SwellDirection: 12.4, SwellHeight: 2.41, SwellPeriod: 17.67,
						WaveDirection: 23.1, WaveHeight: 0.46,
						WindDirection: 131.4, WindSpeed: 10.1,
					},
				),
			)
		},
	)
}

func TestRatingService_GetPositionFromLocation(t *testing.T) {
	t.Run(
		`"It should return "N" for a location of 0 degrees`, func(t *testing.T) {
			setup()
			assert.Equal(t, position.N, ratingService.GetPositionFromLocation(0))
		},
	)

	t.Run(
		`"It should return "E" for a location of 119 degrees`, func(t *testing.T) {
			setup()
			assert.Equal(t, position.E, ratingService.GetPositionFromLocation(119))
		},
	)

	t.Run(
		`"It should return "S" for a location of 219 degrees`, func(t *testing.T) {
			setup()
			assert.Equal(t, position.S, ratingService.GetPositionFromLocation(219))
		},
	)

	t.Run(
		`"It should return "W" for a location of 309 degrees`, func(t *testing.T) {
			setup()
			assert.Equal(t, position.W, ratingService.GetPositionFromLocation(309))
		},
	)
}

func TestRatingService_GetRatingBasedOnWindAndWaveDirections(t *testing.T) {
	t.Run(
		"It should return 1 for a beach with onshore winds", func(t *testing.T) {
			setup()
			assert.Equal(t, 1, ratingService.GetRatingBasedOnWindAndWaveDirections(position.N, position.N))
		},
	)

	t.Run(
		"It should return 3 for a beach with cross winds", func(t *testing.T) {
			setup()
			assert.Equal(t, 3, ratingService.GetRatingBasedOnWindAndWaveDirections(position.E, position.S))
		},
	)

	t.Run(
		"It should return 3 for a beach with offshore winds", func(t *testing.T) {
			setup()
			assert.Equal(t, 5, ratingService.GetRatingBasedOnWindAndWaveDirections(position.S, position.N))
		},
	)
}

func TestRatingService_GetRatingBasedOnSwellPeriod(t *testing.T) {
	t.Run(
		"It should return 1 for a period of 6 seconds", func(t *testing.T) {
			setup()
			assert.Equal(t, 1, ratingService.GetRatingBasedOnSwellPeriod(6))
		},
	)

	t.Run(
		"It should return 2 for a period of 8 seconds", func(t *testing.T) {
			setup()
			assert.Equal(t, 2, ratingService.GetRatingBasedOnSwellPeriod(8))
		},
	)

	t.Run(
		"It should return 4 for a period of 13 seconds", func(t *testing.T) {
			setup()
			assert.Equal(t, 4, ratingService.GetRatingBasedOnSwellPeriod(13))
		},
	)

	t.Run(
		"It should return 5 for a period of 14 seconds", func(t *testing.T) {
			setup()
			assert.Equal(t, 5, ratingService.GetRatingBasedOnSwellPeriod(14))
		},
	)
}

func TestRatingService_GetRatingBasedOnSwellHeight(t *testing.T) {
	t.Run(
		"It should return 1 for a height of 0.2 meters", func(t *testing.T) {
			setup()
			assert.Equal(t, 1, ratingService.GetRatingBasedOnSwellHeight(0.2))
		},
	)

	t.Run(
		"It should return 2 for a height of 0.9 meters", func(t *testing.T) {
			setup()
			assert.Equal(t, 2, ratingService.GetRatingBasedOnSwellHeight(0.9))
		},
	)

	t.Run(
		"It should return 3 for a height of 1.9 meters", func(t *testing.T) {
			setup()
			assert.Equal(t, 3, ratingService.GetRatingBasedOnSwellHeight(1.9))
		},
	)

	t.Run(
		"It should return 5 for a height of 2 meters", func(t *testing.T) {
			setup()
			assert.Equal(t, 5, ratingService.GetRatingBasedOnSwellHeight(2))
		},
	)
}
