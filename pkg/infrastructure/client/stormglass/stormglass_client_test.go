package stormglass

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app/point"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/infrastructure/httpclient/mocked"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/test"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/test/fixture"
)

func TestStormglassClient_FetchPoints(t *testing.T) {
	var (
		mockedHTTPClient *mocked.HTTPClient
		stormglassClient *StormglassClient
	)

	setup := func() {
		mockedHTTPClient = &mocked.HTTPClient{}
		stormglassClient = NewStormglassClient(mockedHTTPClient, "stormglass_url", "stormglass_token")
	}

	t.Run(
		"It should parse and return two points from stormglass response", func(t *testing.T) {
			setup()
			response := test.
				NewResponseBuilder().
				Status(http.StatusOK).
				StrBody(fixture.StormglassTwoValidPointsJSONResponse).
				Build()
			mockedHTTPClient.On("Do", mock.Anything).Return(response, nil)

			points, err := stormglassClient.FetchPoints(100, 100)

			assert.NoError(t, err)
			assert.Equal(
				t, point.Point{
					Time:           time.Date(2020, 04, 26, 00, 00, 00, 00, time.UTC),
					SwellDirection: 64.26, SwellHeight: 0.15, SwellPeriod: 3.89,
					WaveDirection: 23.38, WaveHeight: 0.47,
					WindDirection: 19.45, WindSpeed: 10.1,
				}, points[0],
			)
			assert.Equal(
				t, point.Point{
					Time:           time.Date(2020, 04, 26, 01, 00, 00, 00, time.UTC),
					SwellDirection: 12.4, SwellHeight: 0.21, SwellPeriod: 3.67,
					WaveDirection: 23.1, WaveHeight: 0.46,
					WindDirection: 131.4, WindSpeed: 10.1,
				}, points[1],
			)
			assert.Len(t, points, 2)
		},
	)

	t.Run(
		"It should parse and return the valid point from stormglass response", func(t *testing.T) {
			setup()
			response := test.
				NewResponseBuilder().
				Status(http.StatusOK).
				StrBody(fixture.StormglassOneValidPointJSONResponse).
				Build()
			mockedHTTPClient.On("Do", mock.Anything).Return(response, nil)

			points, err := stormglassClient.FetchPoints(100, 100)

			assert.NoError(t, err)
			assert.Equal(
				t, point.Point{
					Time:           time.Date(2020, 04, 26, 00, 00, 00, 00, time.UTC),
					SwellDirection: 64.26, SwellHeight: 0.15, SwellPeriod: 3.89,
					WaveDirection: 23.38, WaveHeight: 0.47,
					WindDirection: 19.45, WindSpeed: 10.1,
				}, points[0],
			)
			assert.Len(t, points, 1)
		},
	)

	t.Run(
		"It should return an error when stormglass responds with a status different from 200", func(t *testing.T) {
			setup()
			response := test.
				NewResponseBuilder().
				Status(http.StatusTooManyRequests).
				StrBody(fixture.StormglassRateLimitReachedJSONResponse).
				Build()
			mockedHTTPClient.On("Do", mock.Anything).Return(response, nil)

			points, err := stormglassClient.FetchPoints(100, 100)

			assert.Empty(t, points)
			assert.Equal(t, ErrInvalidResponse, err)
		},
	)

	t.Run(
		"It should return an error when stormglass responds with an invalid JSON", func(t *testing.T) {
			setup()
			response := test.
				NewResponseBuilder().
				Status(http.StatusOK).
				StrBody("invalid json").
				Build()
			mockedHTTPClient.On("Do", mock.Anything).Return(response, nil)

			points, err := stormglassClient.FetchPoints(100, 100)

			assert.Error(t, err)
			assert.Empty(t, points)
		},
	)

	t.Run(
		"It should return an error when stormglass request fails", func(t *testing.T) {
			setup()
			mockedHTTPClient.On("Do", mock.Anything).Return(nil, errors.New("some error"))

			points, err := stormglassClient.FetchPoints(100, 100)

			assert.Error(t, err)
			assert.Empty(t, points)
		},
	)
}
