package main

import (
	"log"

	_ "github.com/lib/pq"
)

func main() {
	dsn := initDsn()

	db, err := openDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
