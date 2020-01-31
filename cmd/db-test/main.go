package main

import (
	"log"
	"math/rand"
)

func main() {
	dsn := initDsn()

	db, err := openDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = Insert(db, "hogehoge", rand.Intn(100))
	if err != nil {
		log.Fatal(err)
	}

	err = MultiSelect(db)
	if err != nil {
		log.Fatal(err)
	}

	err = SingleSelect(db)
	if err != nil {
		log.Fatal(err)
	}
}
