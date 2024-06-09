package controller

import (
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
	beaches := []entity.Beach{
		{
			Lat:      1.1,
			Lng:      1.1,
			Name:     "***",
			Position: entity.N,
		},
	}
	timeForecasts, err := controller.beachForecastService.GetBeachForecasts(beaches)
	if err != nil {
		HandleJSON(response, ErrResponseInternalServerError, http.StatusInternalServerError)
		return
	}
	HandleJSON(response, timeForecasts, http.StatusOK)
}
