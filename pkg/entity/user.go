package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name, email, password string) (User, error) {
	now := time.Now()
	uuid := uuid.NewString()
	return User{
		ID:        uuid,
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
