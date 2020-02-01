package main

import (
	"database/sql"
	"fmt"

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

// MultiSelect returns records from the table.
func MultiSelect(db *sql.DB) error {
	stmt := `SELECT * FROM users`
	rows, _ := db.Query(stmt)
	defer rows.Close()
	var users []User
	var user User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Created)
		if err != nil {
			return err
		}
		users = append(users, user)
	}
	for _, user = range users {
		fmt.Println(user)
	}
	return nil
}

// SingleSelect returns a record from the table.
func SingleSelect(db *sql.DB) error {
	stmt := `SELECT * FROM users WHERE age = $1`
	row := db.QueryRow(stmt, 47)
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Created)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	return err
		// } else {
		// 	return err
		// }
		return err
	}
	fmt.Println(user)
	return nil
}

// Delete deletes a record from the table.
func Delete(db *sql.DB) error {
	stmt := `DELETE FROM users WHERE name = $1`
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = db.Exec(stmt, "hogehoge")
	if err != nil {
		return err
	}

	err = tx.Commit()
	return nil
}
