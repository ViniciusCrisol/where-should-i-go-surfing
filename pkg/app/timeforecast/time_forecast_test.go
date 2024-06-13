package timeforecast

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app/point"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity/position"
)

func TestNewBeachForecast(t *testing.T) {
	t.Run(
		"It should return a new BeachForecast with correct values", func(t *testing.T) {
			beach := entity.Beach{
				Lat:      1.1,
				Lng:      1.1,
				Name:     "***",
				Position: position.N,
			}
			point := point.Point{
				Time:           time.Date(2020, 04, 26, 00, 00, 00, 00, time.FixedZone("", 0)),
				SwellDirection: 64.26, SwellHeight: 0.15, SwellPeriod: 3.89,
				WaveDirection: 23.38, WaveHeight: 0.47,
				WindDirection: 19.45, WindSpeed: 10.1,
			}

			beachForecast := NewBeachForecast(beach, point, 5)

			assert.Equal(t, point.Time, beachForecast.Time)
			assert.Equal(t, beach.Lat, beachForecast.Lat)
			assert.Equal(t, beach.Lng, beachForecast.Lng)
			assert.Equal(t, beach.Name, beachForecast.Name)
			assert.Equal(t, beach.Position.String(), beachForecast.Position)
			assert.Equal(t, point.SwellDirection, beachForecast.SwellDirection)
			assert.Equal(t, point.SwellHeight, beachForecast.SwellHeight)
			assert.Equal(t, point.SwellPeriod, beachForecast.SwellPeriod)
			assert.Equal(t, point.WaveDirection, beachForecast.WaveDirection)
			assert.Equal(t, point.WaveHeight, beachForecast.WaveHeight)
			assert.Equal(t, point.WindDirection, beachForecast.WindDirection)
			assert.Equal(t, point.WindSpeed, beachForecast.WindSpeed)
			assert.Equal(t, 5, beachForecast.Rating)
		},
	)
}
