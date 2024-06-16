package test

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func SetupDB() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5433/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	DB = db
}

func ResetDB() {
	if _, err := DB.Exec("TRUNCATE TABLE users;"); err != nil {
		panic(err)
	}
}
