package bcrypt

import (
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

const hashCost = 14

func Hash(plain string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), hashCost)
	if err != nil {
		slog.Error(
			"Failed to hash plain text",
			slog.String("err", err.Error()),
			slog.String("plain_text", plain),
		)
		return "", err
	}
	return string(hashed), err
}

func Compare(plain, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
