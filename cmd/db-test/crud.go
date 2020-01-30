package main

import "database/sql"

// DB is a variable to establish a database connection pool.
var DB *sql.DB

// Insert inserts the record to the table.
func Insert(name string, age int) error {
	stmt := `INSERT INTO users (name, age) VALUES(?, ?)`
	_, err := DB.Exec(stmt, name, age)
	if err != nil {
		return err
	}

	return nil
}
