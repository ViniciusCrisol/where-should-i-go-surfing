package jwt

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var tokenSecretKey []byte

func Init(secretKey string) {
	tokenSecretKey = []byte(secretKey)
}

func SignToken[T any](data T) (string, error) {
	token, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"data": data,
			"exp":  time.Now().Add(time.Hour).Unix(),
		},
	).SignedString(tokenSecretKey)
	if err != nil {
		slog.Error("Failed to sign token", slog.String("err", err.Error()), slog.Any("data", data))
		return "", err
	}
	return token, nil
}

func ParseToken[T any](token string) (*T, bool) {
	jwtToken, err := jwt.Parse(token, func(*jwt.Token) (any, error) { return tokenSecretKey, nil })
	if err != nil {
		return nil, false
	}
	if !jwtToken.Valid {
		return nil, false
	}

	// It's important to cast the data to JSON here because the "data" field within the claims is stored
	// as a generic any type. To convert this generic type back into the original struct type T, we need
	// to serialize it to JSON and then deserialize it into the desired type. This ensures that we
	// correctly preserve the structure and types of the original data, avoiding potential type
	// assertion errors and maintaining the integrity of the data.
	claims, parsed := jwtToken.Claims.(jwt.MapClaims)
	if !parsed {
		slog.Error("Failed to parse claims", slog.String("token", token))
		return nil, false
	}
	buff := new(bytes.Buffer)
	if err = json.NewEncoder(buff).Encode(claims["data"]); err != nil {
		slog.Error("Failed to encode claims", slog.String("token", token))
		return nil, false
	}
	var data T
	if err = json.NewDecoder(buff).Decode(&data); err != nil {
		slog.Error("Failed to decode data", slog.String("token", token))
		return nil, false
	}
	return &data, true
}
