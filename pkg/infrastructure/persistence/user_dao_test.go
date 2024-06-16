package persistence

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/test"
)

func TestUserDAO_Save(t *testing.T) {
	user, _ := entity.NewUser("John Doe", "john.doe@email.com", "123456")
	// PostgreSQL stores timestamps with microsecond precision, so the dates are truncated for test consistency
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
			defer test.ResetDB()

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
			defer test.ResetDB()
			assert.NoError(t, userDAO.Save(user))
			assert.Error(t, userDAO.Save(user))
		},
	)

	t.Run(
		"It should return an error when command fails", func(t *testing.T) {
			setup()
			test.DB.Close()
			assert.Error(t, userDAO.Save(user))
		},
	)
}

func TestUserDAO_FindByEmail(t *testing.T) {}
