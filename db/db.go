package db

import "database/sql"

type DataSource struct {
	Db *sql.DB
}

func NewDataSource(config map[string]interface{}) *DataSource {
	return &DataSource{Db: nil}
}
