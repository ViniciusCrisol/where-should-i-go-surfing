package jwt

import (
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const signKey = "TODO: Change it for an env var!"

func SignToken[T any](data T) (string, error) {
	token, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"data": data,
			"exp":  time.Now().Add(time.Hour).Unix(),
		},
	).SignedString(signKey)
	if err != nil {
		slog.Error("Failed to sign token", slog.String("err", err.Error()), slog.Any("data", data))
		return "", err
	}
	return token, nil
}

func ParseToken[T any](token string) (*T, bool) {
	jwtToken, err := jwt.Parse(token, func(*jwt.Token) (any, error) { return signKey, nil })
	if err != nil {
		return nil, false
	}
	if !jwtToken.Valid {
		return nil, false
	}

	claims, parsed := jwtToken.Claims.(jwt.MapClaims)
	if !parsed {
		slog.Error("Failed to parse claims", slog.String("token", token))
		return nil, false
	}
	data, parsed := claims["data"].(T)
	if !parsed {
		slog.Error("Failed to parse data", slog.String("token", token))
		return nil, false
	}
	return &data, true
}
