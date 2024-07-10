package jwt

import (
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id string) (string, error) {
	token, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  id,
			"exp": time.Now().Add(time.Hour).Unix(),
		},
	).SignedString("TODO: Change it for an env var!")
	if err != nil {
		slog.Error("Failed to sign token", slog.String("err", err.Error()), slog.String("id", id))
		return "", err
	}
	return token, nil
}

func VerifyToken(token string) error {
	return nil
}
