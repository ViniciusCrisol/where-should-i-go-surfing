package controller

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/app"
	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
)

type ErrResponse struct {
	Err string `json:"error"`
}

var (
	ErrResponseInternalServerError  = ErrResponse{"Internal server error"}
	ErrResponseUnsupportedMediaType = ErrResponse{"Unsupported media type, expected application/json"}
)

func HandleErr(response http.ResponseWriter, err error) {
	if errors.Is(err, app.ErrEmailIsAlreadyInUse) {
		HandleJSON(response, ErrResponse{err.Error()}, http.StatusBadRequest)
		return
	}
	if errors.Is(err, entity.ErrInvalidBeachName) {
		HandleJSON(response, ErrResponse{err.Error()}, http.StatusBadRequest)
		return
	}
	if errors.Is(err, entity.ErrInvalidBeachPosition) {
		HandleJSON(response, ErrResponse{err.Error()}, http.StatusBadRequest)
		return
	}
	if errors.Is(err, entity.ErrInvalidUserName) {
		HandleJSON(response, ErrResponse{err.Error()}, http.StatusBadRequest)
		return
	}
	if errors.Is(err, entity.ErrInvalidUserEmail) {
		HandleJSON(response, ErrResponse{err.Error()}, http.StatusBadRequest)
		return
	}
	if errors.Is(err, entity.ErrInvalidUserPassword) {
		HandleJSON(response, ErrResponse{err.Error()}, http.StatusBadRequest)
		return
	}
	HandleJSON(response, ErrResponseInternalServerError, http.StatusInternalServerError)
}

func HandleJSON(response http.ResponseWriter, body any, code int) {
	response.WriteHeader(code)
	response.Header().Set("X-Content-Type-Options", "nosniff")
	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(response).Encode(body); err != nil {
		slog.Error("Failed to encode response ", slog.String("err", err.Error()), slog.Any("body", body))
	}
}
