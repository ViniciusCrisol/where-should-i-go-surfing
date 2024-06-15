package entity

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/helper/bcrypt"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

var (
	ErrInvalidUserName     = errors.New("entity: invalid user name")
	ErrInvalidUserEmail    = errors.New("entity: invalid user email")
	ErrInvalidUserPassword = errors.New("entity: invalid user password")
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
	if len(name) < 3 ||
		len(name) > 256 {
		return User{}, ErrInvalidUserName
	}
	if len(email) < 3 ||
		len(email) > 256 ||
		!emailRegex.MatchString(email) {
		return User{}, ErrInvalidUserEmail
	}
	if len(password) < 6 ||
		len(password) > 256 {
		return User{}, ErrInvalidUserPassword
	}
	hashedPassword, err := bcrypt.Hash(password)
	if err != nil {
		return User{}, err
	}
	now := time.Now()
	uuid := uuid.NewString()
	return User{
		ID:        uuid,
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
