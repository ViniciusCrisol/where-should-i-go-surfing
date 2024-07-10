package entity

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	t.Run(
		"It should return a user", func(t *testing.T) {
			expectedName := "John Doe"
			expectedEmail := "john.doe@email.com"
			expectedPassword := "123456"

			user, err := NewUser(expectedName, expectedEmail, expectedPassword)

			assert.NoError(t, err)
			assert.NotEmpty(t, user.ID)
			assert.Equal(t, expectedName, user.Name)
			assert.Equal(t, expectedEmail, user.Email)
			assert.NotEqual(t, expectedPassword, user.Password)
			assert.NotZero(t, user.CreatedAt)
			assert.NotZero(t, user.UpdatedAt)
		},
	)

	t.Run(
		"It should return an error when the name is too short", func(t *testing.T) {
			_, err := NewUser(strings.Repeat("*", 2), "john.doe@email.com", "123456")
			assert.Equal(t, ErrInvalidUserName, err)
		},
	)

	t.Run(
		"It should return an error when the name is too long", func(t *testing.T) {
			_, err := NewUser(strings.Repeat("*", 257), "john.doe@email.com", "123456")
			assert.Equal(t, ErrInvalidUserName, err)
		},
	)

	t.Run(
		"It should return an error when the email is invalid", func(t *testing.T) {
			_, err := NewUser("John Doe", "john.doe#email.com", "123456")
			assert.Equal(t, ErrInvalidUserEmail, err)
		},
	)

	t.Run(
		"It should return an error when the email is too short", func(t *testing.T) {
			_, err := NewUser("John Doe", strings.Repeat("*", 2), "123456")
			assert.Equal(t, ErrInvalidUserEmail, err)
		},
	)

	t.Run(
		"It should return an error when the email is too long", func(t *testing.T) {
			_, err := NewUser("John Doe", strings.Repeat("*", 257), "123456")
			assert.Equal(t, ErrInvalidUserEmail, err)
		},
	)

	t.Run(
		"It should return an error when the password is too short", func(t *testing.T) {
			_, err := NewUser("John Doe", "john.doe@email.com", strings.Repeat("*", 5))
			assert.Equal(t, ErrInvalidUserPassword, err)
		},
	)

	t.Run(
		"It should return an error when the password is too long", func(t *testing.T) {
			_, err := NewUser("John Doe", "john.doe@email.com", strings.Repeat("*", 257))
			assert.Equal(t, ErrInvalidUserPassword, err)
		},
	)
}

func TestUser_ComparePassword(t *testing.T) {
	user, _ := NewUser("John Doe", "john.doe@email.com", "123456")

	t.Run(
		"It should return true for a correct user password and password combination", func(t *testing.T) {
			assert.True(t, user.ComparePassword("123456"))
		},
	)

	t.Run(
		"It should return false for an incorrect user password and password combination", func(t *testing.T) {
			assert.False(t, user.ComparePassword("1234567"))
		},
	)
}
