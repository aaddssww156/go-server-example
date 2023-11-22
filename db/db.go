package db

import (
	"database/sql"
	"fmt"
	"time"
)

type DB struct {
	DB *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

func Connect(dsn string) (*DB, error) {
	d, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to a database: %e", err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	if err := d.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to a database: %e", err)
	}

	dbConn.DB = d
	return dbConn, nil
}
