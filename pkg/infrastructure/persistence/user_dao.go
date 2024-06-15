package persistence

import (
	"database/sql"
	"errors"
	"log/slog"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
)

const (
	saveUserCommand = `
		INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6);
	`
	findUserByEmail = `
		SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = $1;
	`
)

type UserDAO struct {
	db *sql.DB
}

func NewUserDAO(db *sql.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Save(user entity.User) error {
	if _, err := dao.db.Exec(
		saveUserCommand,
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	); err != nil {
		slog.Error("Failed to execute command", slog.String("err", err.Error()))
		return err
	}
	return nil
}

func (dao *UserDAO) FindByEmail(email string) (entity.User, bool, error) {
	var user entity.User
	if err := dao.db.QueryRow(findUserByEmail, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			slog.Error("Failed to scan row", slog.String("err", err.Error()))
			return entity.User{}, false, err
		}
		return entity.User{}, false, nil
	}
	return user, true, nil
}
