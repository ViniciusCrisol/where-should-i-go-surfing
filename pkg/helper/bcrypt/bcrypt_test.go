package bcrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	t.Run(
		"It should return a valid hash for a given plain text", func(t *testing.T) {
			hashed, err := Hash("123456")
			assert.NoError(t, err)
			assert.NotEmpty(t, hashed)
		},
	)

	t.Run(
		"It should return different hashes for the same plain text due to salting", func(t *testing.T) {
			plain := "123456"
			hashed1, err1 := Hash(plain)
			hashed2, err2 := Hash(plain)
			assert.NoError(t, err1)
			assert.NoError(t, err2)
			assert.NotEqual(t, hashed1, hashed2)
		},
	)
}

func TestCompare(t *testing.T) {
	t.Run(
		"It should return true for a correct plain text and hash combination", func(t *testing.T) {
			plain := "123456"
			hashed, _ := Hash(plain)
			assert.True(t, Compare(plain, hashed))
		},
	)

	t.Run(
		"It should return false for an incorrect plain text and hash combination", func(t *testing.T) {
			hashedText, _ := Hash("123456")
			assert.False(t, Compare("1234567", hashedText))
		},
	)

	t.Run(
		"It should handle invalid hash formats gracefully", func(t *testing.T) {
			assert.False(t, Compare("123456", "invalid hash format"))
		},
	)
}
