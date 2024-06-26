package persistence

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/test"
)

func TestUserDAO_Save(t *testing.T) {
	// PostgreSQL stores timestamps with microsecond precision, so the dates are truncated for test consistency
	user, _ := entity.NewUser("John Doe", "john.doe@email.com", "123456")
	user.CreatedAt = user.CreatedAt.Truncate(time.Microsecond)
	user.UpdatedAt = user.UpdatedAt.Truncate(time.Microsecond)

	var userDAO *UserDAO

	setup := func() {
		test.SetupDB()
		userDAO = NewUserDAO(test.DB)
	}

	t.Run(
		"It should successfully create a new user when email is not in use", func(t *testing.T) {
			setup()

			err := userDAO.Save(user)

			storedUser, _, _ := userDAO.FindByID(user.ID)
			assert.NoError(t, err)
			assert.Equal(t, user.ID, storedUser.ID)
			assert.Equal(t, user.Name, storedUser.Name)
			assert.Equal(t, user.Email, storedUser.Email)
			assert.Equal(t, user.Password, storedUser.Password)
			assert.True(t, user.CreatedAt.Equal(storedUser.CreatedAt))
			assert.True(t, user.UpdatedAt.Equal(storedUser.UpdatedAt))
		},
	)

	t.Run(
		"It should return an error when the user email is already in use", func(t *testing.T) {
			setup()

			assert.NoError(t, userDAO.Save(user))
			assert.Error(t, userDAO.Save(user))
		},
	)

	t.Run(
		"It should return an error when the command fails", func(t *testing.T) {
			setup()
			test.DB.Close()

			assert.Error(t, userDAO.Save(user))
		},
	)
}

func TestUserDAO_FindByID(t *testing.T) {
	// PostgreSQL stores timestamps with microsecond precision, so the dates are truncated for test consistency
	user, _ := entity.NewUser("John Doe", "john.doe@email.com", "123456")
	user.CreatedAt = user.CreatedAt.Truncate(time.Microsecond)
	user.UpdatedAt = user.UpdatedAt.Truncate(time.Microsecond)

	var userDAO *UserDAO

	setup := func() {
		test.SetupDB()
		userDAO = NewUserDAO(test.DB)
	}

	t.Run(
		"It should return a user when the user exists", func(t *testing.T) {
			setup()

			assert.NoError(t, userDAO.Save(user))

			storedUser, found, err := userDAO.FindByID(user.ID)
			assert.NoError(t, err)
			assert.True(t, found)
			assert.Equal(t, user.ID, storedUser.ID)
			assert.Equal(t, user.Name, storedUser.Name)
			assert.Equal(t, user.Email, storedUser.Email)
			assert.Equal(t, user.Password, storedUser.Password)
			assert.True(t, user.CreatedAt.Equal(storedUser.CreatedAt))
			assert.True(t, user.UpdatedAt.Equal(storedUser.UpdatedAt))
		},
	)

	t.Run(
		"It should return false when the user does not exist", func(t *testing.T) {
			setup()

			storedUser, found, err := userDAO.FindByID(user.ID)

			assert.Equal(t, entity.User{}, storedUser)
			assert.False(t, found)
			assert.NoError(t, err)
		},
	)

	t.Run(
		"It should return an error when the query fails", func(t *testing.T) {
			setup()
			test.DB.Close()

			storedUser, found, err := userDAO.FindByID(user.ID)

			assert.Equal(t, entity.User{}, storedUser)
			assert.False(t, found)
			assert.Error(t, err)
		},
	)
}

func TestUserDAO_FindByEmail(t *testing.T) {
	// PostgreSQL stores timestamps with microsecond precision, so the dates are truncated for test consistency
	user, _ := entity.NewUser("John Doe", "john.doe@email.com", "123456")
	user.CreatedAt = user.CreatedAt.Truncate(time.Microsecond)
	user.UpdatedAt = user.UpdatedAt.Truncate(time.Microsecond)

	var userDAO *UserDAO

	setup := func() {
		test.SetupDB()
		userDAO = NewUserDAO(test.DB)
	}

	t.Run(
		"It should return a user when the user exists", func(t *testing.T) {
			setup()

			assert.NoError(t, userDAO.Save(user))

			storedUser, found, err := userDAO.FindByEmail(user.Email)
			assert.NoError(t, err)
			assert.True(t, found)
			assert.Equal(t, user.ID, storedUser.ID)
			assert.Equal(t, user.Name, storedUser.Name)
			assert.Equal(t, user.Email, storedUser.Email)
			assert.Equal(t, user.Password, storedUser.Password)
			assert.True(t, user.CreatedAt.Equal(storedUser.CreatedAt))
			assert.True(t, user.UpdatedAt.Equal(storedUser.UpdatedAt))
		},
	)

	t.Run(
		"It should return false when the user does not exist", func(t *testing.T) {
			setup()

			storedUser, found, err := userDAO.FindByEmail(user.Email)

			assert.Equal(t, entity.User{}, storedUser)
			assert.False(t, found)
			assert.NoError(t, err)
		},
	)

	t.Run(
		"It should return an error when the query fails", func(t *testing.T) {
			setup()
			test.DB.Close()

			storedUser, found, err := userDAO.FindByEmail(user.Email)

			assert.Equal(t, entity.User{}, storedUser)
			assert.False(t, found)
			assert.Error(t, err)
		},
	)
}
