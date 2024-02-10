package adapters

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgresSql() (*sql.DB, error) {
	db, err := sql.Open("pgx", "postgres://dwdarm:asdasdasd@172.17.0.1:5432/shortify")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
