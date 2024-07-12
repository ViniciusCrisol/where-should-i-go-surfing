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
		"It should return an error when the user validation fails", func(t *testing.T) {
			setup()
			assert.Error(t, userService.CreateUser(CreateUserCmd{}))
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

func TestUserService_AuthenticateService(t *testing.T) {
	email := "john.doe@email.com"
	password := "123456"
	user, _ := entity.NewUser("John Doe", email, password)

	var (
		mockedUserDAO *mocked.UserDAO
		userService   *UserService
	)

	setup := func() {
		mockedUserDAO = &mocked.UserDAO{}
		userService = NewUserService(mockedUserDAO)
	}

	t.Run(
		"It should authenticate a user successfully with a valid email/password combination", func(t *testing.T) {
			setup()
			mockedUserDAO.On("FindByEmail", email).Return(user, true, nil)

			token, err := userService.AuthenticateUser(
				AuthenticateUserCmd{
					Email:    email,
					Password: password,
				},
			)

			assert.NoError(t, err)
			assert.NotEmpty(t, token)
		},
	)

	t.Run(
		"It should return an error when the email does not match any user", func(t *testing.T) {
			setup()
			mockedUserDAO.On("FindByEmail", email).Return(entity.User{}, false, nil)

			token, err := userService.AuthenticateUser(
				AuthenticateUserCmd{
					Email:    email,
					Password: password,
				},
			)

			assert.Equal(t, ErrAuthenticationFailed, err)
			assert.Empty(t, token)
		},
	)

	t.Run(
		"It should return an error when the password does not match", func(t *testing.T) {
			setup()
			mockedUserDAO.On("FindByEmail", email).Return(user, true, nil)

			token, err := userService.AuthenticateUser(
				AuthenticateUserCmd{
					Email:    email,
					Password: "1234567",
				},
			)

			assert.Equal(t, ErrAuthenticationFailed, err)
			assert.Empty(t, token)
		},
	)

	t.Run(
		"It should return an error when FindByEmail method fails", func(t *testing.T) {
			setup()
			mockedUserDAO.On("FindByEmail", email).Return(entity.User{}, false, errors.New("some error"))

			token, err := userService.AuthenticateUser(
				AuthenticateUserCmd{
					Email:    email,
					Password: password,
				},
			)

			assert.Error(t, err)
			assert.Empty(t, token)
		},
	)
}
