package controller

import (
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

func (controller *UserController) CreateUser(response http.ResponseWriter, request *http.Request) {}
