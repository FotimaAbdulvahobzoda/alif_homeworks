package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=fotima password=12345678 dbname=coursedb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Connection Error:", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(6)
	db.SetConnMaxLifetime(6 * time.Minute)

	fmt.Println("Before ping:", db.Stats().OpenConnections)
	if err := db.Ping(); err != nil {
		log.Fatal("Connection Error!:", err)
	}

	stats := db.Stats()
	fmt.Println("After ping:", stats.OpenConnections)
	fmt.Println("In use:", stats.InUse)
	fmt.Println("Idle:", stats.Idle)
	fmt.Println("Max open connections:", stats.MaxOpenConnections)

}
