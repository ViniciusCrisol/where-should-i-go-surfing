package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
)

type BeachForecastController struct {
	beachForecastService *app.BeachForecastService
}

func NewBeachForecastController(beachForecastService *app.BeachForecastService) *BeachForecastController {
	return &BeachForecastController{
		beachForecastService: beachForecastService,
	}
}

func (controller *BeachForecastController) GetUserBeachForecasts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	if request.Header.Get("Content-Type") != "application/json" {
		http.Error(response, MessageUnsupportedMediaType, http.StatusUnsupportedMediaType)
		return
	}

	beaches := []entity.Beach{
		{
			Lat:      1.1,
			Lng:      1.1,
			Name:     "***",
			Position: entity.N,
		},
		{
			Lat:      1.11,
			Lng:      1.11,
			Name:     "****",
			Position: entity.S,
		},
	}
	timeForecasts, err := controller.beachForecastService.GetBeachForecasts(beaches)
	if err != nil {
		http.Error(response, MessageInternalServerError, http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(response).Encode(timeForecasts); err != nil {
		slog.Error("Failed to encode response", slog.String("err", err.Error()))
		http.Error(response, MessageInternalServerError, http.StatusInternalServerError)
		return
	}
}
