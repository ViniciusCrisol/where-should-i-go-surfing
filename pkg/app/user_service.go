package app

import (
	"errors"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app/dto/cmd"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
)

var ErrEmailAlreadyInUse = errors.New("service: email already in use")

type UserService struct {
	userDAO UserDAO
}

func NewUserService(userDAO UserDAO) *UserService {
	return &UserService{
		userDAO: userDAO,
	}
}

func (service *UserService) CreateUser(createUserCmd cmd.CreateUserCmd) error {
	user, err := entity.NewUser(
		createUserCmd.Name,
		createUserCmd.Email,
		createUserCmd.Password,
	)
	if err != nil {
		return err
	}
	_, found, err := service.userDAO.FindByEmail(user.Email)
	if err != nil {
		return err
	}
	if found {
		return ErrEmailAlreadyInUse
	}
	return service.userDAO.Save(user)
}
