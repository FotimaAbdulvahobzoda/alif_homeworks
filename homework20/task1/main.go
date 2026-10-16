package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	conn := "host=localhost port=5432 user=fotima password=987654321 dbname=coursedb2 sslmode=disable"

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal("Connection Error!:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Connection Error!:", err)
	}

	result, err := db.Exec(
		`INSERT INTO books (title, author, price, pages) VALUES ($1, $2, $3, $4)`,
		"Harry Potter: The Prizoner of the Azkaban",
		"J.K.Rolling",
		150,
		569,
	)

	if err != nil {
		log.Fatal("Exec Error!:", err)
	}
	fmt.Println(result)

	result1, err := db.Query(
		`SELECT title, author from books`,
	)
	if err != nil {
		log.Fatal("Reading Error!:", err)
	}
	fmt.Println(result1)

	result2, err := db.Exec(`DELETE FROM books WHERE id=$1`, 5)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := result2.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows)

	no := db.QueryRow(
		`INSERT INTO books (title, author, price, pages) VALUES ($1, $2, $3, $4) RETURNING id`,
		"Harry Potter: The Goblet of Fire",
		"J.K.Rolling",
		167,
		723,
	)

	if no != nil {
		log.Fatal(no)
	}

	result4, err := db.Query(`select * from books`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result4)

}
