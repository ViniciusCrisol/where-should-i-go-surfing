package stormglass

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app/dto/point"
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

func (client *StormglassClient) FetchPoints(lat, lng float64) ([]point.Point, error) {
	now := time.Now()
	request, err := client.newFetchPointsRequest(lat, lng, now, now.AddDate(0, 0, 1))
	if err != nil {
		return nil, err
	}
	response, err := client.getFetchPointsResponse(request)
	if err != nil {
		return nil, err
	}
	return client.mapValidPoints(response), nil
}

func (client *StormglassClient) newFetchPointsRequest(lat, lng float64, start, end time.Time) (*http.Request, error) {
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
	query.Set("start", fmt.Sprintf("%d", start.Unix()))
	query.Set("end", fmt.Sprintf("%d", end.Unix()))
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

func (client *StormglassClient) mapValidPoints(response FetchPointsResponse) []point.Point {
	var points []point.Point
	for _, responsePoint := range response.Points {
		if responsePoint.IsValid() {
			points = append(
				points,
				point.Point{
					Time:           responsePoint.Time,
					SwellDirection: responsePoint.SwellDirection["noaa"],
					SwellHeight:    responsePoint.SwellHeight["noaa"],
					SwellPeriod:    responsePoint.SwellPeriod["noaa"],
					WaveDirection:  responsePoint.WaveDirection["noaa"],
					WaveHeight:     responsePoint.WaveHeight["noaa"],
					WindDirection:  responsePoint.WindDirection["noaa"],
					WindSpeed:      responsePoint.WindSpeed["noaa"],
				},
			)
		}
	}
	return points
}
