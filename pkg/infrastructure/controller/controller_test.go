package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
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
