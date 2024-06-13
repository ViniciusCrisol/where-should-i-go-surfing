package app

import (
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app/point"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
)

type StormglassClient interface {
	FetchPoints(lat, lng float64) ([]point.Point, error)
}

type UserDAO interface {
	Save(user entity.User) error
	FindByEmail(email string) (entity.User, bool, error)
}
