package stormglass

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPoint_IsValid(t *testing.T) {
	t.Run(
		"It should return true when the everything is valid", func(t *testing.T) {
			point := Point{
				Time:           time.Now(),
				SwellDirection: map[string]float64{"noaa": math.MaxFloat64},
				SwellHeight:    map[string]float64{"noaa": math.MaxFloat64},
				SwellPeriod:    map[string]float64{"noaa": math.MaxFloat64},
				WaveDirection:  map[string]float64{"noaa": math.MaxFloat64},
				WaveHeight:     map[string]float64{"noaa": math.MaxFloat64},
				WindDirection:  map[string]float64{"noaa": math.MaxFloat64},
				WindSpeed:      map[string]float64{"noaa": math.MaxFloat64},
			}
			assert.Equal(t, true, point.IsValid())
		},
	)

	t.Run(
		"It should return false when the Time is invalid", func(t *testing.T) {
			point := Point{
				SwellDirection: map[string]float64{"noaa": math.MaxFloat64},
				SwellHeight:    map[string]float64{"noaa": math.MaxFloat64},
				SwellPeriod:    map[string]float64{"noaa": math.MaxFloat64},
				WaveDirection:  map[string]float64{"noaa": math.MaxFloat64},
				WaveHeight:     map[string]float64{"noaa": math.MaxFloat64},
				WindDirection:  map[string]float64{"noaa": math.MaxFloat64},
				WindSpeed:      map[string]float64{"noaa": math.MaxFloat64},
			}
			assert.Equal(t, false, point.IsValid())
		},
	)

	t.Run(
		"It should return false when the SwellDirection is invalid", func(t *testing.T) {
			point := Point{
				Time:          time.Now(),
				SwellHeight:   map[string]float64{"noaa": math.MaxFloat64},
				SwellPeriod:   map[string]float64{"noaa": math.MaxFloat64},
				WaveDirection: map[string]float64{"noaa": math.MaxFloat64},
				WaveHeight:    map[string]float64{"noaa": math.MaxFloat64},
				WindDirection: map[string]float64{"noaa": math.MaxFloat64},
				WindSpeed:     map[string]float64{"noaa": math.MaxFloat64},
			}
			assert.Equal(t, false, point.IsValid())
		},
	)

	t.Run(
		"It should return false when the SwellHeight is invalid", func(t *testing.T) {
			point := Point{
				Time:           time.Now(),
				SwellDirection: map[string]float64{"noaa": math.MaxFloat64},
				SwellPeriod:    map[string]float64{"noaa": math.MaxFloat64},
				WaveDirection:  map[string]float64{"noaa": math.MaxFloat64},
				WaveHeight:     map[string]float64{"noaa": math.MaxFloat64},
				WindDirection:  map[string]float64{"noaa": math.MaxFloat64},
				WindSpeed:      map[string]float64{"noaa": math.MaxFloat64},
			}
			assert.Equal(t, false, point.IsValid())
		},
	)

	t.Run(
		"It should return false when the SwellPeriod is invalid", func(t *testing.T) {
			point := Point{
				Time:           time.Now(),
				SwellDirection: map[string]float64{"noaa": math.MaxFloat64},
				SwellHeight:    map[string]float64{"noaa": math.MaxFloat64},
				WaveDirection:  map[string]float64{"noaa": math.MaxFloat64},
				WaveHeight:     map[string]float64{"noaa": math.MaxFloat64},
				WindDirection:  map[string]float64{"noaa": math.MaxFloat64},
				WindSpeed:      map[string]float64{"noaa": math.MaxFloat64},
			}
			assert.Equal(t, false, point.IsValid())
		},
	)

	t.Run(
		"It should return false when the WaveDirection is invalid", func(t *testing.T) {
			point := Point{
				Time:           time.Now(),
				SwellDirection: map[string]float64{"noaa": math.MaxFloat64},
				SwellHeight:    map[string]float64{"noaa": math.MaxFloat64},
				SwellPeriod:    map[string]float64{"noaa": math.MaxFloat64},
				WaveHeight:     map[string]float64{"noaa": math.MaxFloat64},
				WindDirection:  map[string]float64{"noaa": math.MaxFloat64},
				WindSpeed:      map[string]float64{"noaa": math.MaxFloat64},
			}
			assert.Equal(t, false, point.IsValid())
		},
	)

	t.Run(
		"It should return false when the WaveHeight is invalid", func(t *testing.T) {
			point := Point{
				Time:           time.Now(),
				SwellDirection: map[string]float64{"noaa": math.MaxFloat64},
				SwellHeight:    map[string]float64{"noaa": math.MaxFloat64},
				SwellPeriod:    map[string]float64{"noaa": math.MaxFloat64},
				WaveDirection:  map[string]float64{"noaa": math.MaxFloat64},
				WindDirection:  map[string]float64{"noaa": math.MaxFloat64},
				WindSpeed:      map[string]float64{"noaa": math.MaxFloat64},
			}
			assert.Equal(t, false, point.IsValid())
		},
	)

	t.Run(
		"It should return false when the WindDirection is invalid", func(t *testing.T) {
			point := Point{
				Time:           time.Now(),
				SwellDirection: map[string]float64{"noaa": math.MaxFloat64},
				SwellHeight:    map[string]float64{"noaa": math.MaxFloat64},
				SwellPeriod:    map[string]float64{"noaa": math.MaxFloat64},
				WaveDirection:  map[string]float64{"noaa": math.MaxFloat64},
				WaveHeight:     map[string]float64{"noaa": math.MaxFloat64},
				WindSpeed:      map[string]float64{"noaa": math.MaxFloat64},
			}
			assert.Equal(t, false, point.IsValid())
		},
	)

	t.Run(
		"It should return false when the WindSpeed is invalid", func(t *testing.T) {
			point := Point{
				Time:           time.Now(),
				SwellDirection: map[string]float64{"noaa": math.MaxFloat64},
				SwellHeight:    map[string]float64{"noaa": math.MaxFloat64},
				SwellPeriod:    map[string]float64{"noaa": math.MaxFloat64},
				WaveDirection:  map[string]float64{"noaa": math.MaxFloat64},
				WaveHeight:     map[string]float64{"noaa": math.MaxFloat64},
				WindDirection:  map[string]float64{"noaa": math.MaxFloat64},
			}
			assert.Equal(t, false, point.IsValid())
		},
	)
}
