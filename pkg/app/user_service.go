package app

import (
	"errors"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
)

var ErrEmailIsAlreadyInUse = errors.New("service: email already in use")

type UserService struct {
	userDAO UserDAO
}

func NewUserService(userDAO UserDAO) *UserService {
	return &UserService{
		userDAO: userDAO,
	}
}

type CreateUserCmd struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (service *UserService) CreateUser(cmd CreateUserCmd) error {
	user, err := entity.NewUser(cmd.Name, cmd.Email, cmd.Password)
	if err != nil {
		return err
	}
	_, found, err := service.userDAO.FindByEmail(user.Email)
	if err != nil {
		return err
	}
	if found {
		return ErrEmailIsAlreadyInUse
	}
	return service.userDAO.Save(user)
}
