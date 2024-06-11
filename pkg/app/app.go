package app

import "github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app/dto/point"

type StormglassClient interface {
	FetchPoints(lat, lng float64) ([]point.Point, error)
}
