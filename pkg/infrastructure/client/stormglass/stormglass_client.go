package stormglass

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/infrastructure/httpclient"
)

var ErrInvalidResponse = errors.New("stormglass: invalid response")

type StormglassClient struct {
	httpClient       httpclient.HTTPClient
	stormglassURL    string
	stormglassToken  string
	stormglassSource string
	stormglassParams string
}

func NewStormglassClient(httpClient httpclient.HTTPClient, stormglassURL, stormglassToken string) *StormglassClient {
	return &StormglassClient{
		httpClient:       httpClient,
		stormglassURL:    stormglassURL,
		stormglassToken:  stormglassToken,
		stormglassSource: "noaa",
		stormglassParams: strings.Join(
			[]string{
				"swellDirection", "swellHeight", "swellPeriod",
				"waveDirection", "waveHeight",
				"windDirection", "windSpeed",
			}, ",",
		),
	}
}

func (client *StormglassClient) FetchPoints(lat, lng float64) ([]app.Point, error) {
	request, err := client.newFetchPointsRequest(lat, lng)
	if err != nil {
		return nil, err
	}
	response, err := client.getFetchPointsResponse(request)
	if err != nil {
		return nil, err
	}
	return client.mapValidPoints(response), nil
}

func (client *StormglassClient) newFetchPointsRequest(lat, lng float64) (*http.Request, error) {
	fetchPointsURL, err := url.Parse(client.stormglassURL)
	if err != nil {
		slog.Error(
			"Failed to parse url",
			slog.String("err", err.Error()),
			slog.String("url", client.stormglassURL),
		)
		return nil, err
	}
	query := fetchPointsURL.Query()
	query.Set("lat", fmt.Sprintf("%f", lat))
	query.Set("lng", fmt.Sprintf("%f", lng))
	query.Set("source", client.stormglassSource)
	query.Set("params", client.stormglassParams)
	fetchPointsURL.RawQuery = query.Encode()

	fetchPointsRequest, err := http.NewRequest(http.MethodGet, fetchPointsURL.String(), nil)
	if err != nil {
		slog.Error(
			"Failed to create request",
			slog.String("err", err.Error()),
			slog.String("url", fetchPointsURL.String()),
		)
		return nil, err
	}
	fetchPointsRequest.Header.Set("Authorization", client.stormglassToken)
	return fetchPointsRequest, nil
}

func (client *StormglassClient) getFetchPointsResponse(request *http.Request) (FetchPointsResponse, error) {
	httpResponse, err := client.httpClient.Do(request)
	if err != nil {
		slog.Error(
			"Failed to send request",
			slog.String("err", err.Error()),
			slog.String("url", request.URL.String()),
		)
		return FetchPointsResponse{}, err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		slog.Error(
			"Invalid response",
			slog.String("url", request.URL.String()),
			slog.String("status", httpResponse.Status),
		)
		return FetchPointsResponse{}, ErrInvalidResponse
	}

	var fetchPointsResponse FetchPointsResponse
	if err = json.NewDecoder(httpResponse.Body).Decode(&fetchPointsResponse); err != nil {
		slog.Error(
			"Failed to decode response",
			slog.String("err", err.Error()),
			slog.String("url", request.URL.String()),
		)
		return FetchPointsResponse{}, err
	}
	return fetchPointsResponse, nil
}

func (client *StormglassClient) mapValidPoints(response FetchPointsResponse) []app.Point {
	var points []app.Point
	for _, point := range response.Points {
		if point.IsValid() {
			points = append(
				points,
				app.Point{
					Time:           point.Time,
					SwellDirection: point.SwellDirection["noaa"],
					SwellHeight:    point.SwellHeight["noaa"],
					SwellPeriod:    point.SwellPeriod["noaa"],
					WaveDirection:  point.WaveDirection["noaa"],
					WaveHeight:     point.WaveHeight["noaa"],
					WindDirection:  point.WindDirection["noaa"],
					WindSpeed:      point.WindSpeed["noaa"],
				},
			)
		}
	}
	return points
}
