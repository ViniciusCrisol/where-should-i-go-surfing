package jwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

type testData struct {
	ID    int
	Name  string
	Email string
}

var expectedData = testData{
	ID:    1,
	Name:  "John Doe",
	Email: "john.doe@email.com",
}

func setup() {
	tokenSecretKey = []byte("token secret key")
}

func TestSignToken(t *testing.T) {
	t.Run(
		"It should return a valid token for a given data", func(t *testing.T) {
			setup()

			token, err := SignToken(expectedData)
			assert.NoError(t, err)
			assert.NotEmpty(t, token)
		},
	)
}

func TestParseToken(t *testing.T) {
	t.Run(
		"It should return the correct data for a valid token", func(t *testing.T) {
			setup()
			token, _ := SignToken(expectedData)

			data, valid := ParseToken[testData](token)
			assert.True(t, valid)
			assert.Equal(t, expectedData, *data)
		},
	)

	t.Run(
		"It should return false for an invalid token", func(t *testing.T) {
			setup()

			data, valid := ParseToken[testData]("invalid token")
			assert.False(t, valid)
			assert.Nil(t, data)
		},
	)

	t.Run(
		"It should return false for an expired token", func(t *testing.T) {
			setup()
			token, _ := jwt.NewWithClaims(
				jwt.SigningMethodHS256,
				jwt.MapClaims{
					"data": expectedData,
					"exp":  time.Now().Add(-time.Hour).Unix(),
				},
			).SignedString(tokenSecretKey)

			data, valid := ParseToken[testData](token)
			assert.False(t, valid)
			assert.Nil(t, data)
		},
	)
}
