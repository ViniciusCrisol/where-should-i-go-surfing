package controller

import (
	"net/http"
	"time"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity/position"
)

type BeachForecastController struct {
	beachForecastService *app.BeachForecastService
}

func NewBeachForecastController(beachForecastService *app.BeachForecastService) *BeachForecastController {
	return &BeachForecastController{
		beachForecastService: beachForecastService,
	}
}

func (controller *BeachForecastController) GetUserBeachForecasts(response http.ResponseWriter, _ *http.Request) {
	now := time.Now()
	beaches := []entity.Beach{
		{
			ID:        "1",
			Lat:       1.1,
			Lng:       1.1,
			Name:      "Manly",
			Position:  position.N,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	timeForecasts, err := controller.beachForecastService.GetBeachForecasts(beaches)
	if err != nil {
		HandleJSON(response, ErrResponseInternalServerError, http.StatusInternalServerError)
		return
	}
	HandleJSON(response, timeForecasts, http.StatusOK)
}
