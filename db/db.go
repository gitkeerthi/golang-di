package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type DataSource struct {
	Db *sql.DB
}

func NewDataSource(connectionUri string) *DataSource {
	db, err := sql.Open("postgres", connectionUri)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(1 * time.Second)
	db.SetConnMaxLifetime(30 * time.Second)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return &DataSource{Db: db}
}
