package main

import (
	"fmt"
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

	fmt.Println("Execute an insert statement.")
	err = Insert(db, "hogehoge", rand.Intn(100))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Query a multi select statement.")
	err = MultiSelect(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Query a single select statement.")
	err = SingleSelect(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Execute a delete statement.")
	err = Delete(db)
	if err != nil {
		log.Fatal(err)
	}
}
