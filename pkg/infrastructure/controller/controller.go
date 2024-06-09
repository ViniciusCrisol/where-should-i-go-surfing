package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type ErrResponse struct {
	Err string `json:"error"`
}

var (
	ErrResponseInternalServerError  = ErrResponse{"Internal server error"}
	ErrResponseUnsupportedMediaType = ErrResponse{"Unsupported media type, expected application/json"}
)

func HandleJSON(response http.ResponseWriter, body any, status int) {
	response.WriteHeader(status)
	response.Header().Set("X-Content-Type-Options", "nosniff")
	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(response).Encode(body); err != nil {
		slog.Error("Failed to encode response ", slog.String("err", err.Error()), slog.Any("body", body))
	}
}
