package adapters

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgresSql() (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv("POSTGRES_URI"))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
