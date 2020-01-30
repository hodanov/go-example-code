package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Insert inserts the record to the table.
func Insert(db *sql.DB, name string, age int) error {
	stmt := `INSERT INTO users (name, age) VALUES ($1, $2)`
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(stmt, name, age)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return nil
}
