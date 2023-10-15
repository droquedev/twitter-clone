package datastore

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDatabase(dataSourceName string) error {

	var err error

	db, err = sql.Open("postgres", dataSourceName)

	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}
