package position

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
	t.Run(
		`"It should return "N" for North`, func(t *testing.T) {
			assert.Equal(t, "N", N.String())
		},
	)

	t.Run(
		`"It should return "E" for East`, func(t *testing.T) {
			assert.Equal(t, "E", E.String())
		},
	)

	t.Run(
		`"It should return "S" for South`, func(t *testing.T) {
			assert.Equal(t, "S", S.String())
		},
	)

	t.Run(
		`"It should return "W" for West`, func(t *testing.T) {
			assert.Equal(t, "W", W.String())
		},
	)
}
