package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity/position"
)

var (
	ErrInvalidBeachName     = errors.New("entity: invalid beach name")
	ErrInvalidBeachPosition = errors.New("entity: invalid beach position")
)

type Beach struct {
	ID        string
	Lat       float64
	Lng       float64
	Name      string
	Position  position.Position
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBeach(lat, lng float64, name string, position position.Position) (Beach, error) {
	if len(name) < 3 ||
		len(name) > 64 {
		return Beach{}, ErrInvalidBeachName
	}
	if !position.IsValid() {
		return Beach{}, ErrInvalidBeachPosition
	}
	now := time.Now()
	uuid := uuid.NewString()
	return Beach{
		ID:        uuid,
		Lat:       lat,
		Lng:       lng,
		Name:      name,
		Position:  position,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
