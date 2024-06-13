package app

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/test/mocked"
)

func TestUserService_CreateUser(t *testing.T) {
	email := "john.doe@email.com"
	user, _ := entity.NewUser("John Doe", email, "123456")

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
			mockedUserDAO.
				On("FindByEmail", email).
				Return(entity.User{}, false, nil)
			mockedUserDAO.On("Save", user).Return(nil)

			assert.NoError(
				t, userService.CreateUser(
					CreateUserCmd{
						Name:     user.Name,
						Email:    user.Email,
						Password: user.Password,
					},
				),
			)
		},
	)

	t.Run(
		"It should return an error when the user email is already in use", func(t *testing.T) {
			setup()
			mockedUserDAO.
				On("FindByEmail", email).
				Return(entity.User{}, true, nil)

			assert.Equal(
				t, ErrEmailIsAlreadyInUse, userService.CreateUser(
					CreateUserCmd{
						Name:     user.Name,
						Email:    user.Email,
						Password: user.Password,
					},
				),
			)
		},
	)

	t.Run(
		"It should return an error when FindByEmail method fails", func(t *testing.T) {
			setup()
			mockedUserDAO.
				On("FindByEmail", email).
				Return(entity.User{}, false, errors.New("some error"))

			assert.Error(
				t, userService.CreateUser(
					CreateUserCmd{
						Name:     user.Name,
						Email:    user.Email,
						Password: user.Password,
					},
				),
			)
		},
	)

	t.Run(
		"It should return an error when Save method fails", func(t *testing.T) {
			setup()
			mockedUserDAO.
				On("FindByEmail", email).
				Return(entity.User{}, false, nil)
			mockedUserDAO.On("Save", user).Return(errors.New("some error"))

			assert.Error(
				t, userService.CreateUser(
					CreateUserCmd{
						Name:     user.Name,
						Email:    user.Email,
						Password: user.Password,
					},
				),
			)
		},
	)
}
