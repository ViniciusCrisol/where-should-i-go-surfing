package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app"
)

type UserController struct {
	userService *app.UserService
}

func NewUserController(userService *app.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (controller *UserController) CreateUser(response http.ResponseWriter, request *http.Request) {
	if request.Header.Get("Content-Type") != "application/json" {
		HandleJSON(response, ErrResponseUnsupportedMediaType, http.StatusUnsupportedMediaType)
		return
	}
	var cmd app.CreateUserCmd
	if err := json.NewDecoder(request.Body).Decode(&cmd); err != nil {
		HandleJSON(response, ErrResponseUnsupportedMediaType, http.StatusUnsupportedMediaType)
		return
	}
	if err := controller.userService.CreateUser(cmd); err != nil {
		HandleErr(response, err)
		return
	}
	response.WriteHeader(http.StatusCreated)
}
