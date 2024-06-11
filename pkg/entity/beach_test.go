package entity

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity/position"
)

func TestNewBeach(t *testing.T) {
	t.Run(
		"It should return a beach", func(t *testing.T) {
			expectedLat := 100.11
			expectedLng := 100.11
			expectedName := "***"
			expectedPosition := position.N

			beach, err := NewBeach(expectedLat, expectedLng, expectedName, expectedPosition)

			assert.NoError(t, err)
			assert.Equal(t, expectedLat, beach.Lat)
			assert.Equal(t, expectedLng, beach.Lng)
			assert.Equal(t, expectedName, beach.Name)
			assert.Equal(t, expectedPosition, beach.Position)
		},
	)

	t.Run(
		"It should return an error when the name is too short", func(t *testing.T) {
			_, err := NewBeach(100, 100, strings.Repeat("*", 2), position.N)
			assert.Equal(t, ErrInvalidBeachName, err)
		},
	)

	t.Run(
		"It should return an error when the name is too long", func(t *testing.T) {
			_, err := NewBeach(100, 100, strings.Repeat("*", 65), position.N)
			assert.Equal(t, ErrInvalidBeachName, err)
		},
	)

	t.Run(
		"It should return an error when the position is invalid", func(t *testing.T) {
			_, err := NewBeach(100, 100, "Ana", "*")
			assert.Equal(t, ErrInvalidBeachPosition, err)
		},
	)
}
