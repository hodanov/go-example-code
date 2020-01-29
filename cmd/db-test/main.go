package main

import (
	"log"
)

func main() {
	dsn := initDsn()

	db, err := openDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
