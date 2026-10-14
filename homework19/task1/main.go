package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=fotima password=12345678 dbname=coursedb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Connection Error:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Connection Error!", err)
	}

	fmt.Println("Connection to the database is successful!")
}
