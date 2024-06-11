package entity

import (
	"errors"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity/position"
)

var (
	ErrInvalidName     = errors.New("entity: invalid name")
	ErrInvalidPosition = errors.New("entity: invalid position")
)

type Beach struct {
	Lat      float64
	Lng      float64
	Name     string
	Position position.Position
}

func NewBeach(lat, lng float64, name string, position position.Position) (Beach, error) {
	if len(name) < 3 ||
		len(name) > 64 {
		return Beach{}, ErrInvalidName
	}
	if !position.IsValid() {
		return Beach{}, ErrInvalidPosition
	}
	return Beach{
		Lat:      lat,
		Lng:      lng,
		Name:     name,
		Position: position,
	}, nil
}
