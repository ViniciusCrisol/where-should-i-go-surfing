package logger

import (
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Init()
	slog.Info("INFO")

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	assert.Contains(t, string(out), `"msg":"INFO"`)
	assert.Contains(t, string(out), `"level":"INFO"`)
	assert.Contains(t, string(out), "where-should-i-go-surfing/pkg/helper/logger/logger_test.go:18")
}