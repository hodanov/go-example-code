package main

import (
	"database/sql"
	"fmt"
)

func initDsn() string {
	var dsn string

	var (
		// server   = os.Getenv("POSTGRES_SERVER")
		// port     = os.Getenv("POSTGRES_PORT")
		// user     = os.Getenv("POSTGRES_USER")
		// password = os.Getenv("POSTGRES_PASSWORD")
		// database = os.Getenv("POSTGRES_DB")
		server   = "psql"
		port     = "5432"
		user     = "test"
		password = "test"
		database = "test"
	)
	// dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, server, port, database)
	dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, server, port, database)

	return dsn
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
