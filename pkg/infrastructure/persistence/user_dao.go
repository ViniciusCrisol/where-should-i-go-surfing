package persistence

import (
	"database/sql"
	"errors"
	"log/slog"

	"github.com/ViniciusCrisol/where-should-i-go-surfing/pkg/entity"
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
		"INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6);",
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	); err != nil {
		slog.Error("Failed to execute command", slog.String("error", err.Error()), slog.Any("user", user))
		return err
	}
	return nil
}

func (dao *UserDAO) FindByID(id string) (entity.User, bool, error) {
	row := dao.db.QueryRow(
		"SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1;",
		id,
	)
	if err := row.Err(); err != nil {
		slog.Error("Failed to execute query", slog.String("err", err.Error()), slog.String("user_id", id))
		return entity.User{}, false, err
	}
	var user entity.User
	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			slog.Error("Failed to scan row", slog.String("err", err.Error()), slog.String("user_id", id))
			return entity.User{}, false, err
		}
		return entity.User{}, false, nil
	}
	return user, true, nil
}

func (dao *UserDAO) FindByEmail(email string) (entity.User, bool, error) {
	row := dao.db.QueryRow(
		"SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = $1;",
		email,
	)
	if err := row.Err(); err != nil {
		slog.Error("Failed to execute query", slog.String("err", err.Error()), slog.String("user_email", email))
		return entity.User{}, false, err
	}
	var user entity.User
	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			slog.Error("Failed to scan row", slog.String("err", err.Error()), slog.String("user_email", email))
			return entity.User{}, false, err
		}
		return entity.User{}, false, nil
	}
	return user, true, nil
}
