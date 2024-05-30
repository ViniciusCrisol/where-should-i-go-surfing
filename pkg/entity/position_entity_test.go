package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPosition_IsValid(t *testing.T) {
	t.Run(
		"It should return true when position is valid", func(t *testing.T) {
			position := Position("N")
			assert.True(t, position.IsValid())
			position = "E"
			assert.True(t, position.IsValid())
			position = "S"
			assert.True(t, position.IsValid())
			position = "W"
			assert.True(t, position.IsValid())
		},
	)

	t.Run(
		"It should return false when position is invalid", func(t *testing.T) {
			position := Position("*")
			assert.False(t, position.IsValid())
		},
	)
}

func TestPosition_String(t *testing.T) {
	assert.Equal(t, "N", N.String())
	assert.Equal(t, "E", E.String())
	assert.Equal(t, "S", S.String())
	assert.Equal(t, "W", W.String())
}
