// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocked

import (
	"github.com/stretchr/testify/mock"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app/point"
)

// StormglassClient is an autogenerated mock type for the StormglassClient type
type StormglassClient struct {
	mock.Mock
}

// FetchPoints provides a mock function with given fields: lat, lng
func (_m *StormglassClient) FetchPoints(lat float64, lng float64) ([]point.Point, error) {
	ret := _m.Called(lat, lng)

	if len(ret) == 0 {
		panic("no return value specified for FetchPoints")
	}

	var r0 []point.Point
	var r1 error
	if rf, ok := ret.Get(0).(func(float64, float64) ([]point.Point, error)); ok {
		return rf(lat, lng)
	}
	if rf, ok := ret.Get(0).(func(float64, float64) []point.Point); ok {
		r0 = rf(lat, lng)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]point.Point)
		}
	}

	if rf, ok := ret.Get(1).(func(float64, float64) error); ok {
		r1 = rf(lat, lng)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStormglassClient creates a new instance of StormglassClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStormglassClient(
	t interface {
		mock.TestingT
		Cleanup(func())
	},
) *StormglassClient {
	mock := &StormglassClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
