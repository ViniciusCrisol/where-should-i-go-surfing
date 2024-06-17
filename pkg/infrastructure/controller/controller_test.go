package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
)

func GetStrResponseBody(response *http.Response) string {
	body, _ := io.ReadAll(response.Body)
	return string(body)
}

func GetErrResponseResponseBody(response *http.Response) ErrResponse {
	var errResponse ErrResponse
	json.NewDecoder(response.Body).Decode(&errResponse)
	return errResponse
}

func TestHandleErr(t *testing.T) {
	t.Run(
		"It should return BadRequest for ErrEmailIsAlreadyInUse", func(t *testing.T) {
			response := httptest.NewRecorder()

			HandleErr(response, app.ErrEmailIsAlreadyInUse)

			assert.Equal(t, http.StatusBadRequest, response.Code)
			assert.JSONEq(t, fmt.Sprintf(`{"error":"%s"}`, app.ErrEmailIsAlreadyInUse), response.Body.String())
		},
	)

	t.Run(
		"It should return BadRequest for ErrInvalidBeachName", func(t *testing.T) {
			response := httptest.NewRecorder()

			HandleErr(response, entity.ErrInvalidBeachName)

			assert.Equal(t, http.StatusBadRequest, response.Code)
			assert.JSONEq(t, fmt.Sprintf(`{"error":"%s"}`, entity.ErrInvalidBeachName), response.Body.String())
		},
	)

	t.Run(
		"It should return BadRequest for ErrInvalidBeachPosition", func(t *testing.T) {
			response := httptest.NewRecorder()

			HandleErr(response, entity.ErrInvalidBeachPosition)

			assert.Equal(t, http.StatusBadRequest, response.Code)
			assert.JSONEq(t, fmt.Sprintf(`{"error":"%s"}`, entity.ErrInvalidBeachPosition), response.Body.String())
		},
	)

	t.Run(
		"It should return BadRequest for ErrInvalidUserName", func(t *testing.T) {
			response := httptest.NewRecorder()

			HandleErr(response, entity.ErrInvalidUserName)

			assert.Equal(t, http.StatusBadRequest, response.Code)
			assert.JSONEq(t, fmt.Sprintf(`{"error":"%s"}`, entity.ErrInvalidUserName), response.Body.String())
		},
	)

	t.Run(
		"It should return BadRequest for ErrInvalidUserEmail", func(t *testing.T) {
			response := httptest.NewRecorder()

			HandleErr(response, entity.ErrInvalidUserEmail)

			assert.Equal(t, http.StatusBadRequest, response.Code)
			assert.JSONEq(t, fmt.Sprintf(`{"error":"%s"}`, entity.ErrInvalidUserEmail), response.Body.String())
		},
	)

	t.Run(
		"It should return BadRequest for ErrInvalidUserPassword", func(t *testing.T) {
			response := httptest.NewRecorder()

			HandleErr(response, entity.ErrInvalidUserPassword)

			assert.Equal(t, http.StatusBadRequest, response.Code)
			assert.JSONEq(t, fmt.Sprintf(`{"error":"%s"}`, entity.ErrInvalidUserPassword), response.Body.String())
		},
	)

	t.Run(
		"It should return InternalServerError for unknown errors", func(t *testing.T) {
			response := httptest.NewRecorder()

			HandleErr(response, errors.New("some error"))

			assert.Equal(t, http.StatusInternalServerError, response.Code)
			assert.JSONEq(t, fmt.Sprintf(`{"error":"%s"}`, ErrResponseInternalServerError.Err), response.Body.String())
		},
	)
}

func TestHandleJSON(t *testing.T) {
	t.Run(
		"It should set the correct response values", func(t *testing.T) {
			response := httptest.NewRecorder()

			HandleJSON(response, map[string]bool{"ok": true}, http.StatusOK)

			assert.Equal(t, http.StatusOK, response.Code)
			assert.JSONEq(t, `{"ok": true}`, response.Body.String())
			assert.Equal(t, "nosniff", response.Header().Get("X-Content-Type-Options"))
			assert.Equal(t, "application/json; charset=utf-8", response.Header().Get("Content-Type"))
		},
	)
}
