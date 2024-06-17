package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app/point"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app/timeforecast"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity/position"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/infrastructure/client/stormglass"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/infrastructure/httpclient/mocked"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/test"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/test/fixture"
)

func TestBeachForecastController_GetBeachForecasts(t *testing.T) {
	now := time.Now()
	beach := entity.Beach{
		ID:        "1",
		Lat:       1.1,
		Lng:       1.1,
		Name:      "Manly",
		Position:  position.N,
		CreatedAt: now,
		UpdatedAt: now,
	}
	point1 := point.Point{
		Time:           time.Date(2020, 04, 26, 00, 00, 00, 00, time.UTC),
		SwellDirection: 64.26, SwellHeight: 0.15, SwellPeriod: 3.89,
		WaveDirection: 23.38, WaveHeight: 0.47,
		WindDirection: 19.45, WindSpeed: 10.1,
	}
	point2 := point.Point{
		Time:           time.Date(2020, 04, 26, 01, 00, 00, 00, time.UTC),
		SwellDirection: 12.4, SwellHeight: 0.21, SwellPeriod: 3.67,
		WaveDirection: 23.1, WaveHeight: 0.46,
		WindDirection: 131.4, WindSpeed: 10.1,
	}

	var (
		httpServer       *httptest.Server
		mockedHTTPClient *mocked.HTTPClient
	)

	setup := func() {
		mockedHTTPClient = &mocked.HTTPClient{}
		stormglassClient := stormglass.NewStormglassClient(
			mockedHTTPClient,
			"stormglass_url",
			"stormglass_token",
		)
		beachForecastService := app.NewBeachForecastService(stormglassClient)
		beachForecastController := NewBeachForecastController(beachForecastService)
		httpServer = httptest.NewServer(http.HandlerFunc(beachForecastController.GetUserBeachForecasts))
	}

	t.Run(
		"It should return multiple forecasts with different times for the user beach", func(t *testing.T) {
			setup()
			defer httpServer.Close()
			stormglassResponse := test.
				NewResponseBuilder().
				Status(http.StatusOK).
				StrBody(fixture.StormglassTwoValidPointsJSONResponse).
				Build()
			mockedHTTPClient.On("Do", mock.Anything).Return(stormglassResponse, nil)
			expectedJSONResponse, _ := json.Marshal(
				timeforecast.NewTimeForecastsBuilder().
					BeachForecast(point1.Time, timeforecast.NewBeachForecast(beach, point1, 1)).
					BeachForecast(point2.Time, timeforecast.NewBeachForecast(beach, point2, 2)).
					Build(),
			)

			response, _ := http.Get(httpServer.URL)

			assert.Equal(t, http.StatusOK, response.StatusCode)
			assert.JSONEq(t, string(expectedJSONResponse), GetStrResponseBody(response))
		},
	)

	t.Run(
		"It should return an error when stormglass client returns an error", func(t *testing.T) {
			setup()
			defer httpServer.Close()
			mockedHTTPClient.On("Do", mock.Anything).Return(nil, errors.New("some error"))

			response, _ := http.Get(httpServer.URL)

			assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
			assert.Equal(t, ErrResponseInternalServerError, GetErrResponseResponseBody(response))
		},
	)
}
