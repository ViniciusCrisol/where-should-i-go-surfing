package persistence

import (
	"database/sql"

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
	return nil
}

func (dao *UserDAO) FindByEmail(email string) (entity.User, bool, error) {
	return entity.User{}, false, nil
}
