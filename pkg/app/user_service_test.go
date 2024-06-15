package app

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/infrastructure/persistence/mocked"
)

func TestUserService_CreateUser(t *testing.T) {
	name := "John Doe"
	email := "john.doe@email.com"
	password := "123456"

	var (
		mockedUserDAO *mocked.UserDAO
		userService   *UserService
	)

	setup := func() {
		mockedUserDAO = &mocked.UserDAO{}
		userService = NewUserService(mockedUserDAO)
	}

	t.Run(
		"It should successfully create a new user when email is not in use", func(t *testing.T) {
			setup()
			mockedUserDAO.On("FindByEmail", email).Return(entity.User{}, false, nil)
			mockedUserDAO.On("Save", mock.AnythingOfType("entity.User")).Return(nil)

			assert.NoError(
				t, userService.CreateUser(
					CreateUserCmd{
						Name:     name,
						Email:    email,
						Password: password,
					},
				),
			)
		},
	)

	t.Run(
		"It should return an error when the user email is already in use", func(t *testing.T) {
			setup()
			mockedUserDAO.On("FindByEmail", email).Return(entity.User{}, true, nil)

			assert.Equal(
				t, ErrEmailIsAlreadyInUse, userService.CreateUser(
					CreateUserCmd{
						Name:     name,
						Email:    email,
						Password: password,
					},
				),
			)
		},
	)

	t.Run(
		"It should return an error when FindByEmail method fails", func(t *testing.T) {
			setup()
			mockedUserDAO.On("FindByEmail", email).Return(entity.User{}, false, errors.New("some error"))

			assert.Error(
				t, userService.CreateUser(
					CreateUserCmd{
						Name:     name,
						Email:    email,
						Password: password,
					},
				),
			)
		},
	)

	t.Run(
		"It should return an error when Save method fails", func(t *testing.T) {
			setup()
			mockedUserDAO.On("FindByEmail", email).Return(entity.User{}, false, nil)
			mockedUserDAO.On("Save", mock.AnythingOfType("entity.User")).Return(errors.New("some error"))

			assert.Error(
				t, userService.CreateUser(
					CreateUserCmd{
						Name:     name,
						Email:    email,
						Password: password,
					},
				),
			)
		},
	)
}
