package app

import (
	"errors"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/helper/jwt"
)

var (
	ErrEmailIsAlreadyInUse  = errors.New("service: email already in use")
	ErrAuthenticationFailed = errors.New("service: authentication failed")
)

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
	_, found, err := service.userDAO.FindByEmail(cmd.Email)
	if err != nil {
		return err
	}
	if found {
		return ErrEmailIsAlreadyInUse
	}
	return service.userDAO.Save(user)
}

type AuthenticateUserCmd struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (service *UserService) AuthenticateUser(cmd AuthenticateUserCmd) (string, error) {
	user, found, err := service.userDAO.FindByEmail(cmd.Email)
	if err != nil {
		return "", err
	}
	if !found {
		return "", ErrAuthenticationFailed
	}
	if !user.ComparePassword(cmd.Password) {
		return "", ErrAuthenticationFailed
	}
	return jwt.SignToken(user.ID)
}
