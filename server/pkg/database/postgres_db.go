package database

import (
	"database/sql"
	"twitter-clone/server/config"

	_ "github.com/lib/pq"
)

func InitDatabase(config *config.Config) (*sql.DB, error) {

	db, err := sql.Open("postgres", config.DB_URI)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db, nil
}
