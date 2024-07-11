package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/infrastructure/persistence"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/test"
)

func TestUserController_CreateUser(t *testing.T) {
	requestBody := `{"name": "John Doe", "email": "john.doe@email.com", "password": "123456"}`

	var httpServer *httptest.Server

	setup := func() {
		test.SetupDB()
		userDAO := persistence.NewUserDAO(test.DB)
		userService := app.NewUserService(userDAO)
		userController := NewUserController(userService)
		httpServer = httptest.NewServer(http.HandlerFunc(userController.CreateUser))
	}

	t.Run(
		"It should create a new user successfully when the email is not in use", func(t *testing.T) {
			setup()

			response, _ := http.Post(httpServer.URL, "application/json", strings.NewReader(requestBody))

			assert.Equal(t, http.StatusCreated, response.StatusCode)
			assert.Empty(t, GetStrResponseBody(response))
		},
	)

	t.Run(
		"It should return an error when the email is already in use", func(t *testing.T) {
			setup()

			http.Post(httpServer.URL, "application/json", strings.NewReader(requestBody))
			response, _ := http.Post(httpServer.URL, "application/json", strings.NewReader(requestBody))

			assert.Equal(t, http.StatusBadRequest, response.StatusCode)
			assert.Equal(t, ErrResponse{app.ErrEmailIsAlreadyInUse.Error()}, GetErrResponseResponseBody(response))
		},
	)

	t.Run(
		"It should return an error if the database connection fails", func(t *testing.T) {
			setup()
			test.DB.Close()

			response, _ := http.Post(httpServer.URL, "application/json", strings.NewReader(requestBody))

			assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
			assert.Equal(t, ErrResponseInternalServerError, GetErrResponseResponseBody(response))
		},
	)

	t.Run(
		"It should return an error when the Content-Type header is not application/json", func(t *testing.T) {
			setup()

			response, _ := http.Post(httpServer.URL, "application/xml", strings.NewReader(requestBody))

			assert.Equal(t, http.StatusUnsupportedMediaType, response.StatusCode)
			assert.Equal(t, ErrResponseUnsupportedMediaType, GetErrResponseResponseBody(response))
		},
	)

	t.Run(
		"It should return an error if the request body unmarshal fails", func(t *testing.T) {
			setup()

			response, _ := http.Post(
				httpServer.URL,
				"application/json",
				strings.NewReader(
					`
						<?xml version="1.0" encoding="UTF-8" ?>
						 <root>
						     <name>John Doe</name>
						     <email>john.doe@email.com</email>
						     <password>123456</password>
						 </root>
					`,
				),
			)

			assert.Equal(t, http.StatusUnsupportedMediaType, response.StatusCode)
			assert.Equal(t, ErrResponseUnsupportedMediaType, GetErrResponseResponseBody(response))
		},
	)
}

func TestUserController_AuthenticateUser(t *testing.T) {
	requestBody := `{"email": "john.doe@email.com", "password": "123456"}`

	var httpServer *httptest.Server

	setup := func() {
		test.SetupDB()
		userDAO := persistence.NewUserDAO(test.DB)
		userService := app.NewUserService(userDAO)
		userController := NewUserController(userService)
		httpServer = httptest.NewServer(http.HandlerFunc(userController.AuthenticateUser))

		userService.CreateUser(app.CreateUserCmd{Name: "John Doe", Email: "john.doe@email.com", Password: "123456"})
	}

	t.Run(
		"It should authenticate a user successfully with a valid email/password combination", func(t *testing.T) {
			setup()

			response, _ := http.Post(httpServer.URL, "application/json", strings.NewReader(requestBody))

			assert.Equal(t, http.StatusOK, response.StatusCode)
			assert.NotEmpty(t, GetStrResponseBody(response))
		},
	)

	t.Run(
		"It should return an error when the email does not match any user", func(t *testing.T) {
			setup()

			response, _ := http.Post(
				httpServer.URL,
				"application/json",
				strings.NewReader(`{"email": "john_doe@email.com", "password": "123456"}`),
			)

			assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
			assert.Equal(t, ErrResponse{app.ErrAuthenticationFailed.Error()}, GetErrResponseResponseBody(response))
		},
	)

	t.Run(
		"It should return an error when the password does not match", func(t *testing.T) {
			setup()

			response, _ := http.Post(
				httpServer.URL,
				"application/json",
				strings.NewReader(`{"email": "john.doe@email.com", "password": "1234567"}`),
			)

			assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
			assert.Equal(t, ErrResponse{app.ErrAuthenticationFailed.Error()}, GetErrResponseResponseBody(response))
		},
	)

	t.Run(
		"It should return an error if the database connection fails", func(t *testing.T) {
			setup()
			test.DB.Close()

			response, _ := http.Post(httpServer.URL, "application/json", strings.NewReader(requestBody))

			assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
			assert.Equal(t, ErrResponseInternalServerError, GetErrResponseResponseBody(response))
		},
	)

	t.Run(
		"It should return an error when the Content-Type header is not application/json", func(t *testing.T) {
			setup()

			response, _ := http.Post(httpServer.URL, "application/xml", strings.NewReader(requestBody))

			assert.Equal(t, http.StatusUnsupportedMediaType, response.StatusCode)
			assert.Equal(t, ErrResponseUnsupportedMediaType, GetErrResponseResponseBody(response))
		},
	)

	t.Run(
		"It should return an error if the request body unmarshal fails", func(t *testing.T) {
			setup()

			response, _ := http.Post(
				httpServer.URL,
				"application/json",
				strings.NewReader(
					`
						<?xml version="1.0" encoding="UTF-8" ?>
						 <root>
						     <email>john.doe@email.com</email>
						     <password>123456</password>
						 </root>
					`,
				),
			)

			assert.Equal(t, http.StatusUnsupportedMediaType, response.StatusCode)
			assert.Equal(t, ErrResponseUnsupportedMediaType, GetErrResponseResponseBody(response))
		},
	)
}
